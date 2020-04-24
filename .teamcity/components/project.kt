import jetbrains.buildServer.configs.kotlin.v2019_2.Project
import java.io.File

object AzureRM : Project({
    vcsRoot(providerRepository)

    var environment = "public" // TODO: configurable
    var customParallelism = mapOf(
        "containers" to serviceDetails("Containers", 5, 5),
        "compute" to serviceDetails("Compute", 5, 4)
    )
    File("services.txt").forEachLine { l ->
        var serviceName = l
        var service = customParallelism.get(serviceName)
        var build = buildConfigurationForService(environment, serviceName, service)
        buildType(build)
    }
})