import jetbrains.buildServer.configs.kotlin.v2019_2.BuildType
import jetbrains.buildServer.configs.kotlin.v2019_2.buildSteps.script

fun buildConfigurationForService(environment : String, serviceName: String): BuildType {
    return BuildType {
        // TC needs a consistent ID for dynamically generated packages
        // luckily Golang packages are valid, so we're good to reuse that
        id("AZURE_SERVICE_%s_%s".format(environment.toUpperCase(), serviceName.toUpperCase()))
        name = "Service - %s".format(serviceName)

        vcs {
            root(providerRepository)
            cleanCheckout = true
        }

        steps {
            script {
                name = "Validation"
                scriptContent = "echo \"Hello World\""
            }
        }

        failureConditions {
            errorMessage = true
        }
    }
}