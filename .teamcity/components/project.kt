import jetbrains.buildServer.configs.kotlin.v2019_2.Project

object AzureRM : Project({
    vcsRoot(providerRepository)
    buildType(providerBuild)
})