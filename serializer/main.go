package main

import (
	"fmt"
	"serializer/flatstruct"
	"serializer/fmtdatetime"
	"serializer/int2float"
	"serializer/nestedstruct"
	"serializer/numberofstring"
	"serializer/oacstruct"
)

func main() {
	flatstruct.FlatStructSerializer()
	fmt.Println("/*-----------------------------------*/")
	nestedstruct.NestedStructSerializer()
	fmt.Println("/*-----------------------------------*/")
	oacstruct.OACStructSerializer()
	fmt.Println("/*-----------------------------------*/")
	numberofstring.NumberOfStringFieldDeserializer()
	fmt.Println("/*-----------------------------------*/")
	int2float.Int2FloatDeserializer()
	fmt.Println("/*-----------------------------------*/")
	fmtdatetime.FmtDatetimeSerializer()
	fmt.Println("/*-----------------------------------*/")
}
