# Ascii Art Web - Stylize

Ascii Art Web is a web application that allows users to generate ASCII art from text and various banner styles.

## Description

Ascii Art Web Stylize is a web application that allows users to generate ASCII art from text and various banner styles. The application provides a user-friendly interface for users to input their desired text and select a banner style. The application then processes the input and generates the corresponding ASCII art. This project utilizes css in style.css file that been created to make a universal style for the website. 

## Authors

- Hawra AlFaraj(halfaraj)
- Alaa AlMahroos(aalmahroo)
- Yusuf Hubail(yhubail)

## Usage: how to run

1. Clone the repository.
2. Navigate to the root directory of the project.
3. Run the command `go run main.go`.
4. Open your browser and navigate to "http://localhost:8080".

# Note: 
To interupt the server, go back to terminal & Press Ctrl + C

## Implementation details: algorithm

The algorithm for generating ASCII art is straightforward, utilizing HTML, CSS and server backend development as well as error handling. The application reads the content of the selected banner / font file as in the original ascii art project. It then iterates over each character of the input text in the text area and replaces it with the corresponding character from the banner file. The resulting string is then displayed as the ASCII art in the ascii art page.