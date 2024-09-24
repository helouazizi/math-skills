# math-skills

This Go application performs statistical calculations on a dataset contained in a text file. <br> It computes the average, median, variance, and standard <br>deviation of a series of numbers.

# Features
Reads numerical data from a file named data.txt.<br>
Computes:<br>
Average<br>
Median<br>
Variance
Standard Deviation<br>


//////////////////////////////////////////

# Running the Program
Run the application using the following command:

`go run . data.txt`

# Output
The program will output the following statistical measures:

Average<br>
Median<br>
Variance<br>
Standard Deviatio<br>

////////////////////////////////////////////////////

# Example output:

average:  5.4<br>
median:  5.5<br>
variance:  2.3<br>
standard_deviation:  1.5<br>



/////////////////////////////////////////////////////

# Functions
The application utilizes functions defined in the functions package, which includes:<br>

ReadFile: Reads the contents of a specified file.<br>
Parse_Float: Converts the file contents into a slice of floats.<br>
Average: Calculates the average of the float slice.<br>
Median: Computes the median value of the float slice.<br>
Variance: Determines the variance based on the float slice and average.<br>
Standar_Devision: Calculates the standard deviation from the variance.<br>


## Contributing

Contributions are welcome! If you have suggestions for improvements or new features, feel free to open an issue or submit a pull request.

