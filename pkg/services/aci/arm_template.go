package aci

// nolint: lll
var armTemplateBytes = []byte(`
	{
		"$schema": "https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
		"contentVersion": "1.0.0.0",
		"parameters": {
			"location": {
				"type": "string"
			},
			"name": {
				"type": "string",
				"metadata": {
					"description": "Name for the container group"
				},
				"defaultValue": "acilinuxpublicipcontainergroup"
			},
			"image": {
				"type": "string",
				"metadata": {
					"description": "Container image to deploy. Should be of the form accountName/imagename:tag for images stored in Docker Hub or a fully qualified URI for a private registry like the Azure Container Registry."
				},
				"defaultValue": "microsoft/aci-helloworld"
			},
			"port": {
				"type": "int",
				"metadata": {
					"description": "Port to open on the container and the public IP address."
				},
				"defaultValue": "80"
			},
			"cpuCores": {
				"type": "int",
				"metadata": {
					"description": "The number of CPU cores to allocate to the container. Must be an integer."
				},
				"defaultValue": "1.0"
			},
			"memoryInGb": {
				"type": "string",
				"metadata": {
					"description": "The amount of memory to allocate to the container in gigabytes."
				},
				"defaultValue": "1.5"
			},
			"tags": {
				"type": "object"
			}
		},
		"variables": {},
		"resources": [
			{
				"name": "[parameters('name')]",
				"type": "Microsoft.ContainerInstance/containerGroups",
				"apiVersion": "2017-08-01-preview",
				"location": "[parameters('location')]",
				"properties": {
					"containers": [
						{
							"name": "[parameters('name')]",
							"properties": {
								"image": "[parameters('image')]",
								"ports": [
									{
										"port": "[parameters('port')]" 
									}
								],
								"resources": {
									"requests": {
										"cpu": "[parameters('cpuCores')]",
										"memoryInGb": "[float(parameters('memoryInGb'))]"
									}
								}
							}
						}
					],
					"osType": "Linux",
					"ipAddress": {
						"type": "Public",
						"ports": [
							{
								"protocol": "tcp",
								"port": "[parameters('port')]"
							}
						]
					}
				},
				"tags": "[parameters('tags')]"
			}
		],
		"outputs": {
			"containerIPv4Address":{
				"type": "string",
				"value": "[reference(resourceId('Microsoft.ContainerInstance/containerGroups/', parameters('name'))).ipAddress.ip]"
			}
		}
	}
`)
