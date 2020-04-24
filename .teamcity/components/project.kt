import jetbrains.buildServer.configs.kotlin.v2019_2.Project
import java.io.File

object AzureRM : Project({
    vcsRoot(providerRepository)

    var environment = "public" // TODO: configurable
    File("services.txt").forEachLine { l ->
        var serviceName = l

        var build = buildConfigurationForService(environment, serviceName)
        buildType(build)
    }
})