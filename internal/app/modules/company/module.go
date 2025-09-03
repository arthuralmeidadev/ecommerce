package company

import "ecommerce/pkg/deuterium"

func Module() *deuterium.Module  {
   return &deuterium.Module{
       Name: "Company",
       Controller: controller(),
   } 
}
