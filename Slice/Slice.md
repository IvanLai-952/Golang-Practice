# Slice

## Basic Description

Slice is a data structure that describes a portion of an array by specifying the starting index and the length of the portion.

It is just like an array having an index value and length, but the size of the slice is resized they are not in fixed-size just like an array. Internally, slice and an array are connected with each other, a slice is a reference to an underlying array.

## Components of Slice

- Pointer: The pointer is used to points to the first element of the array that is accessible through the slice. Here, it is not necessary that the pointed element is the first element of the array.  
- Length: The length is the total number of elements present in the array.  
- Capacity: The capacity represents the maximum size upto which it can expand.  

**Compare the different result between copy() and append()**

## Getting Started

To run the examples provided in this directory, follow these steps:

1. Clone the Repository:
```bash
git clone https://github.com/IvanLai-952/Golang-Practice.git
cd Golang-Practice/Slice
```

2. Run the Program:
Execute the main.go file:
```golang
go run main.go 1 2 3 4 5
```