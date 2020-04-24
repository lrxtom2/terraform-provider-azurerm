import jetbrains.buildServer.configs.kotlin.v2019_2.BuildType
import jetbrains.buildServer.configs.kotlin.v2019_2.buildSteps.script

object providerBuild : BuildType({
    name = "Azure Provider Example"

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
})