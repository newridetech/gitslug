package gitslug

import (
    // things here do not have to be cryptographically secure
    // our goal is to provide unique names
    "crypto/md5"
    "fmt"
    "regexp"

    "github.com/gosimple/slug"
)

func SlugifyGitRepositoryName(repoName string) string {
    urlRegexp := regexp.MustCompile("^git@github.com:([\\w]+)\\/([\\w\\.]+)\\.git$")
    matches := urlRegexp.FindStringSubmatch(repoName)

    vendor := slug.Make(matches[1])
    project := slug.Make(matches[2])

    sum := md5.Sum([]byte(repoName))

    return fmt.Sprintf("%s-%s-%x", project, vendor, sum)
}
