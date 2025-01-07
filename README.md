# URL Shortener Service

## Overview

This project is a **URL shortener service** that allows users to shorten URLs and redirect them when visited. The service is built with **Go (Golang)** for the backend, using **PostgreSQL** as the database for storing URL mappings. The project is hosted on **Heroku** with a containerized deployment, using a **Docker-based CI/CD pipeline** for continuous integration and deployment.

---

## Tech Stack

- **Backend**: Go (Golang)
- **Database**: PostgreSQL (Hosted on Heroku)
- **Containerization**: Docker
- **Cloud Provider**: Heroku (for hosting the application and PostgreSQL database)
- **CI/CD**: GitHub Actions (to automate testing, building, and deployment)
- **Analytics**: Built-in analytics to track usage of the shortened URLs

---

## Features

- **URL Shortening**: Shorten long URLs into shorter, more manageable links.
- **Redirection**: Redirect users to the original URL when the shortened URL is accessed.
- **Analytics**: Track the usage of shortened URLs, including the number of times a URL is accessed.
- **Containerized Deployment**: The application is built and deployed in a Docker container, ensuring consistency across environments.
- **Continuous Integration & Deployment (CI/CD)**: Automated testing, building, and deployment via GitHub Actions.

---

## How It Works

1. **URL Shortening**: When a user sends a POST request to the `/shorten` endpoint with a long URL, the service generates a short URL and stores it in the database.
2. **Redirection**: When a user accesses a shortened URL, the service fetches the original URL from the database and redirects the user to it.
3. **Analytics**: The service tracks the number of accesses for each shortened URL, providing insights into the URL's popularity.

---

## How to Set Up

### Prerequisites

Before running or deploying the project, ensure you have the following:

- **Heroku account**: To deploy the application and PostgreSQL database.
- **Docker**: To containerize the app for local development and deployment.
- **GitHub account**: For CI/CD automation with GitHub Actions.

### Steps to Run Locally

1. **Clone the repository:**

```bash
git clone https://github.com/your-username/url-shortener.git
cd url-shortener
```

2. **Set up the PostgreSQL database:**

- You will need a **PostgreSQL** database. You can use the **PostgreSQL addon** on Heroku or run it locally using **Docker**.
- If using **Heroku**, the database URL will be set as an environment variable.

3. **Set up environment variables:** Create a `.env` file with the following variables:

```env
DATABASE_URL=<your_database_url>
PORT=8080
```

4. **Run the application locally:**

```bash
go run main.go
```

Your app will now be running on `http://localhost:8080`.

## CI/CD with GitHub Actions

The project uses **GitHub Actions** for CI/CD. The workflow is defined in the `.github/workflows/ci.yml` file and does the following:

- **Runs tests**: Every time you push to the `main` branch or create a pull request, tests are run using Docker and Go.
- **Builds Docker image**: After the tests pass, a Docker image is built from the source code.
- **Pushes Docker image**: The built Docker image is pushed to DockerHub.
- **Deploys to Heroku**: The image is then pushed to Heroku's container registry, and the application is redeployed.

---

## Example CI/CD Pipeline (`.github/workflows/ci.yml`)

```yaml
name: CI/CD Pipeline

on:
  push:
    branches:
      - main
  pull_request:
    branches:
      - main

jobs:
  build-and-deploy:
    runs-on: ubuntu-latest

    env:
      DOCKER_USERNAME: ${{ secrets.DOCKER_USERNAME }}
      DOCKER_PASSWORD: ${{ secrets.DOCKER_PASSWORD }}
      HEROKU_API_KEY: ${{ secrets.HEROKU_API_KEY }}
      HEROKU_APP_NAME: ${{ secrets.HEROKU_APP_NAME }}
      IMAGE_TAG: ${{ github.sha }}  # Use commit SHA as the tag

    steps:
    - name: Checkout code
      uses: actions/checkout@v3

    - name: Run tests with Go container
      uses: docker://golang:1.23
      with:
        args: sh -c "go test ./..."

    - name: Set Heroku container stack
      run: |
        curl -n -X PATCH https://api.heroku.com/apps/${{ env.HEROKU_APP_NAME }} \
          -H "Content-Type: application/json" \
          -H "Accept: application/vnd.heroku+json; version=3" \
          -H "Authorization: Bearer $HEROKU_API_KEY" \
          -d '{"build_stack":"container"}'

    - name: Set up Docker
      uses: docker/setup-buildx-action@v2

    - name: Cache Docker layers
      uses: actions/cache@v2
      with:
        path: ~/.cache/docker
        key: ${{ runner.os }}-docker-${{ github.sha }}
        restore-keys: |
          ${{ runner.os }}-docker-

    - name: Log in to DockerHub
      run: |
        echo $DOCKER_PASSWORD | docker login -u $DOCKER_USERNAME --password-stdin

    - name: Build Docker image
      run: docker build -t $DOCKER_USERNAME/url-shortener:${{ env.IMAGE_TAG }} .

    - name: Push Docker image to DockerHub
      run: docker push $DOCKER_USERNAME/url-shortener:${{ env.IMAGE_TAG }}

    - name: Install Heroku CLI
      run: |
        curl https://cli-assets.heroku.com/install.sh | sh

    - name: Log in to Heroku container registry
      run: |
        echo $HEROKU_API_KEY | docker login --username=_ --password-stdin registry.heroku.com

    - name: Deploy to Heroku
      run: |
        docker tag $DOCKER_USERNAME/url-shortener:${{ env.IMAGE_TAG }} registry.heroku.com/${{ env.HEROKU_APP_NAME }}/web
        docker push registry.heroku.com/${{ env.HEROKU_APP_NAME }}/web
        heroku container:release web --app ${{ env.HEROKU_APP_NAME }}
```

### Heroku Setup

1. **Set up the Heroku application:**

- Create a new Heroku app if you haven't already.
- Add the PostgreSQL addon to your Heroku app.

2. **Deploy the app:** The CI/CD pipeline will automatically deploy the Docker container to Heroku upon pushing to the `main` branch.

---

## Learning Objectives

This project is designed to help me deepen my understanding of **DevOps**, **Backend Development**, and **Cloud Infrastructure**. Here's how it helps me learn:

- **Infrastructure as Code**: Using GitHub Actions and Docker to automate the build and deployment process, I’m embracing the concept of Infrastructure as Code.
- **Containerization**: By Dockerizing the application, I ensure it runs consistently across different environments, whether locally or in production.
- **CI/CD Pipelines**: Automating the testing, building, and deployment of the app reduces human error and accelerates the development cycle.
- **Cloud Deployment**: Hosting the app on Heroku allows me to focus on development while Heroku handles the scaling and maintenance of the infrastructure.
- **Backend Development**: Developing the URL shortener service has enhanced my skills in RESTful API development, database management, and analytics.

---

## License

This project is licensed under the **MIT License** - see the LICENSE file for details.
