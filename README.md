## 🧐 About 

This project aims to dockerize the Ascii-Art-WEB application, which allows users to generate ASCII art banners using different styles. By following the instructions below, you will learn how to containerize the application using Docker while adhering to good coding practices and Dockerfile best practices.

### Objectives

The objectives for this project are as follows:
- Create at least one Dockerfile.
- Build an image based on the Dockerfile.
- Run a container using the image.

### Usage
After successfully building the Docker image, you can create and run a Docker container using the image. Use the following command:

`docker container run -p 8080:8080 ascii-go`

Additional Docker Commands:

Here are some additional Docker commands that might be useful:

- Build the Docker image:
`docker image build -t ascii-go .`

- Check the status of running containers:
`docker ps -a`

-  Remove a specific container (replace CONTAINER_ID with the actual ID):
`docker rm CONTAINER_ID`

- View the list of Docker images:
`docker images`

Accessing the Application

Once the Docker container is running, you can access the Ascii-Art-WEB-Stylize application by opening a web browser and navigating to http://localhost:8080.


## ✍️ Authors 

Sabyrbek Madiar [msabyrbe]
Aiana Oserbaeva   [oayana]



5. Open your web browser and go to http://localhost:8080 to access the application.

## Implementation Details
### ASCII Art Generation Algorithm

The ASCII art generation process follows these steps:
◊
1. **ASCII ARTSYLES**:
   - Each character is represented by an 8-line pattern in the art style.
   - Characters are separated by new lines (\n).

2. **User Input Processing**:
   - The application takes user input, which includes the text and the chosen banner style.

3. **Loading Art Styles**:
   - The chosen art style's ASCII art template is loaded from the corresponding text file in the "artstyles" directory.

4. **Text to ASCII Art Conversion**:
   - The input text is split into lines, where each line represents a line of ASCII art.
   - For each line of text:
     - The line is split into individual characters.
     - For each character:
       - The character's ASCII value is checked to ensure it falls within the Basic Latin character range (ASCII values 32 to 126).
       - If the character is within this range, its corresponding pattern from the banner style's ASCII art template is retrieved.
       - The patterns of all characters in the line are concatenated to form the ASCII art representation for that line.
       - The process is repeated for each of the 8 lines in the character's pattern.
       - The application maps each character to its corresponding pattern in the banner style, resulting in the ASCII art representation for each line of text.

5. **Rendering the Result**:
   - The generated ASCII art, along with the original input text and chosen banner style, is rendered on the result page for the user to view.


### Web Application Framework

Ascii-Art-Web is built on the Go programming language and the net/http package. It consists of the following main components:

- **Main Page Handler**: Renders the main page where users input text and select a banner style.

- **ASCII Art Handler**: Processes user input, generates ASCII art, and renders the result page.

- **Not Found Handler**: Handles cases where users access non-existent routes.

## Extensible and Scalable

The application is designed to be extensible, allowing for easy addition of new banner styles by creating corresponding text files in the "banners" directory. 
- To add custome Banner File make sure:
  - Banner Format

    - Each character has a height of 8 lines.
    - Characters are separated by a new line \n.




 