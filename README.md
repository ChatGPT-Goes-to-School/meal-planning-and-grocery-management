# Meal Planning and Grocery Management Application

This repository contains the source code for the Meal Planning and Grocery Management application. The application aims to provide users with a convenient way to plan their meals, create grocery lists, and manage their food inventory.

## Features

- **Meal Planning:** Users can create and manage meal plans for breakfast, lunch, and dinner. They can add recipes to their meal plans and specify the desired dates for each meal.

- **Grocery Management:** Users can create and manage grocery lists to keep track of the items they need to purchase. They can add items to the list, mark them as purchased, and remove them when no longer needed.

- **Recipe Integration:** The application integrates with external recipe APIs and databases, allowing users to search for recipes and add them directly to their meal plans or grocery lists.

- **Ingredient Information:** Users can access detailed information about ingredients, including nutritional data, allergen information, and cooking instructions.

- **Inventory Tracking:** The application provides an inventory management feature, allowing users to keep track of the ingredients they have at home. They can mark items as consumed when used in a meal, and the application will automatically update the inventory.

## Technologies Used

- **Backend:** The backend is built using Golang with the Gin framework, providing a robust and efficient foundation for the application.

- **External APIs:** The application integrates with various third-party APIs for recipe data, ingredient information, and nutritional data. These APIs provide a rich source of data to enhance the user experience.

## Getting Started

To set up the application locally, follow these steps:

1. Clone the repository: `git clone https://github.com/your-username/meal-planning-grocery-management.git`

2. Install the necessary dependencies for the backend by running:
  ```go mod download```

3. Start the development server using:
  ```go run main.go```

4. Access the application in your web browser at `http://localhost:8080`.

## CI/CD Workflow

The Meal Planning and Grocery Management application follows a CI/CD workflow using GitHub Actions for continuous integration and deployment. The workflow is triggered when changes are pushed to the repository, ensuring that the application is built, tested, and deployed automatically.

The CI/CD workflow includes the following steps:

1. **Build:** The workflow starts by building the application's Docker image using the Dockerfile. This ensures that the application is packaged and ready for deployment.

2. **Test:** Automated tests are executed to ensure the functionality and integrity of the application. This includes unit tests, integration tests, and any other relevant tests specific to the modules.

3. **Containerization:** The Docker image is pushed to a container registry (e.g., Amazon ECR) to make it available for deployment.

4. **Deployment:** The application is deployed to AWS EKS (Elastic Kubernetes Service) using Kubernetes manifests. The manifests define the desired state of the application and ensure that it is deployed consistently across different environments.

5. **Configuration:** The necessary AWS resources, such as RDS and Elasticache, are set up and configured for the application. Environment variables and secrets are managed using AWS Secrets Manager to ensure secure and scalable deployment.

6. **Monitoring:** CloudWatch is configured to monitor the health and performance of the deployed application. This includes collecting and analyzing logs, setting up alerts, and visualizing metrics to ensure the application is running smoothly.

The CI/CD workflow ensures that any changes made to the repository are automatically tested and deployed in a controlled and consistent manner. This enables faster development cycles and reduces the risk of introducing bugs or issues into the application.

For more details on the CI/CD pipeline configuration and workflow, please refer to the GitHub Actions workflow files in the repository.

## Cloud Architecture

The Meal Planning and Grocery Management application is designed to be deployed in a cloud environment, leveraging various AWS services for scalability, reliability, and performance. The following diagram illustrates the high-level cloud architecture of the application:

![Cloud Architecture](path/to/cloud_architecture_diagram.png)

The key components of the architecture include:

- **AWS EKS**: Elastic Kubernetes Service is used as the container orchestration platform to manage the deployment and scaling of microservices.

- **AWS RDS**: Amazon RDS provides a managed relational database service for storing data related to meal planning and grocery management.

- **Amazon ElastiCache**: ElastiCache is utilized as an in-memory caching service to cache frequently accessed data, improving application performance and reducing database load.

- **Amazon CloudWatch**: CloudWatch is configured to monitor the health and performance of the application, collecting logs and metrics to enable efficient troubleshooting and proactive monitoring.

- **AWS Secrets Manager**: Secrets Manager is used to securely store and manage environment variables and sensitive information, ensuring proper access control and encryption.

- **Elastic Load Balancer (ELB)**: ELB is employed for load balancing incoming traffic to ensure high availability and distribute requests across multiple instances of the application.

The cloud architecture is designed to provide scalability, fault tolerance, and flexibility for the Meal Planning and Grocery Management application. By leveraging AWS services, the application can handle varying workloads, ensure data reliability, and deliver a seamless user experience.

For a more detailed understanding of the cloud architecture, including specific configurations and deployment patterns, refer to the architecture documentation or consult the relevant architecture diagrams.

## Documentation

For detailed documentation on the API endpoints and usage, refer to the [API Documentation](api-docs.md) file.

## License

This project is licensed under the [MIT License](LICENSE).
