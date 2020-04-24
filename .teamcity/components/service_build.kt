import jetbrains.buildServer.configs.kotlin.v2019_2.BuildType
import jetbrains.buildServer.configs.kotlin.v2019_2.buildSteps.script
import jetbrains.buildServer.configs.kotlin.v2019_2.triggers.ScheduleTrigger
import jetbrains.buildServer.configs.kotlin.v2019_2.triggers.schedule

fun buildConfigurationForService(azureEnv : String, serviceName: String, parallelism : Int): BuildType {
    return BuildType {
        // TC needs a consistent ID for dynamically generated packages
        // luckily Golang packages are valid, so we're good to reuse that
        id("AZURE_SERVICE_%s_%s".format(azureEnv.toUpperCase(), serviceName.toUpperCase()))
        name = "Service - %s".format(serviceName)

        vcs {
            root(providerRepository)
            cleanCheckout = true
        }

        steps {
            script {
                name = "Configure Go Version"
                scriptContent = "goenv install -s \$(goenv local) && goenv rehash"
            }

            script {
                name = "Run Tests"
                scriptContent = "go test -v ./azurerm/internal/services/%s -timeout=%TIMEOUT% -run=%TEST_PREFIX% -json".format(serviceName)
            }
        }

        failureConditions {
            errorMessage = true
        }

        params {
            text("PARALLELISM", "%d".format(parallelism))
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
                    hour = 1 // TODO: does this want to be staggered?
                    timezone = "SERVER"
                }
            }
        }
    }
}