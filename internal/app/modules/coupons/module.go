package coupons

import "ecommerce/pkg/deuterium"

func Module() *deuterium.Module  {
   return &deuterium.Module{
       Name: "Coupons",
       Controller: controller(),
   } 
}
