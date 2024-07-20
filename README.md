
## Prerequisites
 - [Microsoft Azure Account](https://azure.microsoft.com/en-us/free)
 - [GoLang](https://golang.org/doc/install)
 - [Resty](https://github.com/go-resty/resty) and [Godontenv](https://github.com/joho/godotenv)
 - Azure Text Analytics key (Congetive Service) and Endpoint

## Setting Up Azure Cognitive Services

#### 1- Sign in to the [Azure](https://azure.microsoft.com/en-us/get-started/azure-portal) Portal.

#### 2- Click on Create a resource in the left sidebar.
#### 3- Search for  Text Analytics and select it from the list.
#### 4- Click Create.
#### 5-Fill in the required details
#### 6-Click Review and create; then click Create to deploy the resource
## Getting Endpoint and API Key
- Go to your resource in the Azure Portal.
- In the left sidebar select Keys and Endpoint.
- Copy the KEY 1 or KEY 2
- Copy the Endpoint URL.(https://your-resource-name.cognitiveservices.azure.com/)



## Setup Project Lab 

### Setup
- Clone the Repository

```http
https://github.com/ahmedshefo/Azure-text-analyzer

```

- Install Dependencies

```http
go get github.com/go-resty/resty/v2
go get github.com/joho/godotenv

```
- Create a .env File
In the root directory of the project, create a file named .env and write in it:
```http
API_KEY=your_azure_text_analytics_api_key
ENDPOINT=https://your-resource-name.cognitiveservices.azure.com/

```
- Update and Run the Project
 To run the project, use:

```http
go run main.go

```
## Sentiment Response Format

#### The output will be formatted as follows:
- #### For Positive: "It seems to radiate positivity"
- #### For Neutral: "The sentiment is balanced"
- #### For Negative: "It seems to carry a touch of negativity"

## Join us
- #### If you have suggestions ; feel free to fork the repository and submit a pull request.


