import jetbrains.buildServer.configs.kotlin.v2019_2.BuildType
import jetbrains.buildServer.configs.kotlin.v2019_2.buildFeatures.golang
import jetbrains.buildServer.configs.kotlin.v2019_2.buildSteps.script
import jetbrains.buildServer.configs.kotlin.v2019_2.triggers.schedule

class serviceDetails(displayName : String, parallelism: Int, startHour: Int) {
    var displayName = displayName
    var parallelism = parallelism
    var startHour = startHour
}

fun buildConfigurationForService(azureEnv : String, packageName: String,service: serviceDetails?): BuildType {
    return BuildType {
        // TC needs a consistent ID for dynamically generated packages
        // luckily Golang packages are valid, so we're good to reuse that
        id("AZURE_SERVICE_%s_%s".format(azureEnv.toUpperCase(), packageName.toUpperCase()))

        if (service != null) {
            name = "Service - %s".format(service.displayName)
        } else {
            name = "Service - %s".format(packageName)
        }

        vcs {
            root(providerRepository)
            cleanCheckout = true
        }

        steps {
            script {
                name = "Configure Go Version"
                scriptContent = "goenv install -s \$(goenv local) && goenv rehash"
            }

            var servicePath = "./azurerm/internal/services/%s/...".format(packageName)
            script {
                name = "Run Tests"
                scriptContent = "go test -v $servicePath -timeout=%TIMEOUT% -run=%TEST_PREFIX% -json"
            }
        }

        failureConditions {
            errorMessage = true
        }

        features {
            golang {
                testFormat = "json"
            }
        }

        params {
            if (service != null) {
                text("PARALLELISM", "%d".format(service.parallelism))
            } else {
                text("PARALLELISM", "30")
            }
            text("TEST_PREFIX", "TestAcc")
            text("TIMEOUT", "12h")
            text("env.TF_ACC", "1")
            text("env.TF_SCHEMA_PANIC_ON_ERROR", "1")
        }

        triggers {
            schedule {
                enabled = false
                type = "schedulingTrigger"
                branchFilter = "+:refs/heads/master"
                schedulingPolicy = daily {

                    if (service != null) {
                        hour = service.startHour
                    } else {
                        hour = 1
                    }

                    timezone = "SERVER"
                }
            }
        }
    }
}