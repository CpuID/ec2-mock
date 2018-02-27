package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Server struct {
	Port uint64
	Tags InstanceTags
}

func (s *Server) Start() {
	log.Printf("Listening on port %d\n", s.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", s.Port), s.NewMux()); err != nil {
		log.Fatalf("Error creating http server: %+v\n", err)
	}
	log.Printf("Exiting...\n")
}

func (s *Server) NewMux() *http.ServeMux {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", s.rootHandler)
	return mux
}

func (s *Server) rootHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		if req.FormValue("Action") == "DescribeTags" {
			log.Printf("POST / (Action DescribeTags)\n")
			filters := make(map[string]string)
			// 10 will do for now
			for i := 1; i <= 10; i++ {
				filter_name := fmt.Sprintf("Filter.%d.Name", i)
				filter_value := fmt.Sprintf("Filter.%d.Value.1", i)
				if req.FormValue(filter_name) != "" && req.FormValue(filter_value) != "" {
					filters[req.FormValue(filter_name)] = req.FormValue(filter_value)
				}
			}
			response := `<DescribeTagsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>7a62c49f-347e-4fc4-9331-6e8e</requestId>
   <tagSet>
`
			//fmt.Printf("%+v\n", filters)
			if len(filters) > 0 {
				if filters["resource-type"] == "instance" && filters["resource-id"] != "" {
					// Return instance ID specific tags (if any exist)
					if _, ok := s.Tags[filters["resource-id"]]; ok == true {
						for _, v := range s.Tags[filters["resource-id"]] {
							response = fmt.Sprintf(`%s      <item>
         <resourceId>%s</resourceId>
         <resourceType>instance</resourceType>
         <key>%s</key>
         <value>%s</value>
      </item>
`, response, filters["resource-id"], v.Key, v.Value)
						}
					}
				} else {
					io.WriteString(w, "invalid Filter.* field specified")
					return
				}
			} else {
				// Return everything defined, no filters
				for k1, v1 := range s.Tags {
					for _, v2 := range v1 {
						response = fmt.Sprintf(`%s      <item>
         <resourceId>%s</resourceId>
         <resourceType>instance</resourceType>
         <key>%s</key>
         <value>%s</value>
      </item>
`, response, k1, v2.Key, v2.Value)
					}
				}
			}
			response = fmt.Sprintf("%s%s", response, `   </tagSet>
</DescribeTagsResponse>`)
			io.WriteString(w, response)
		} else {
			io.WriteString(w, "invalid Action specified")
		}
	} else {
		io.WriteString(w, "invalid HTTP method specified")
	}
}
