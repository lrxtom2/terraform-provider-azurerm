import jetbrains.buildServer.configs.kotlin.v2019_2.Project
import java.io.File

object AzureRM : Project({
    vcsRoot(providerRepository)

    var environment = "public" // TODO: configurable
    var customParallelism = mapOf(
        "containers" to 5,
        "compute" to 7
    )
    File("services.txt").forEachLine { l ->
        var serviceName = l
        var paralellism = customParallelism.getOrDefault(serviceName, 30)
        var build = buildConfigurationForService(environment, serviceName, paralellism)
        buildType(build)
    }
})