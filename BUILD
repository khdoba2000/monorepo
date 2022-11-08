load("@bazel_gazelle//:def.bzl", "gazelle")

#gazelle:prefix monorepo
#gazelle:go_naming_convention import_alias
gazelle(
    name = "gazelle",
    gazelle = "//:gazelle_binary",
)

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
        "-build_file_proto_mode=disable_global",
    ],
    command = "update-repos",
)

load("@bazel_gazelle//:def.bzl", "DEFAULT_LANGUAGES", "gazelle_binary")

gazelle_binary(
    name = "gazelle_binary",
    languages = DEFAULT_LANGUAGES + ["@golink//gazelle/go_link:go_default_library"],
    visibility = ["//visibility:public"],
)
