# S3 Cache Pull

A bitrise step to download your cache from a s3 bucket using custom keys with fallback.

Should be used with [S3 Cache Push](https://github.com/alephao/bitrise-step-s3-cache-push)

### Inputs

<table>
    <thead>
        <tr>
            <th>Input</th>
            <th>Description</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>
                <b>cache_aws_access_key_id</b>
            </td>
            <td>
                Your aws access key id
            </td>
        </tr>
        <tr>
            <td>
                <b>cache_aws_secret_access_key</b>
            </td>
            <td>
                Your aws secret access key
            </td>
        </tr>
        <tr>
            <td>
                <b>cache_aws_region</b>
            </td>
            <td>
                The region of your S3 bucket. E.g.: <tt>us-east-1</tt>
            </td>
        </tr>
        <tr>
            <td>
                <b>cache_bucket_name</b>
            </td>
            <td>
                The name of your S3 bucket. E.g.: <tt>mybucket</tt>
            </td>
        </tr>
        <tr>
            <td>
                <b>cache_restore_keys</b>
            </td>
            <td>
                <span>The list of keys to restore the cache. E.g.:</span>
                <pre>
carthage-$BRANCH_NAME-$COMMIT_SHA
carthage-$BRANCH_NAME
carthage-
                </pre>
            </td>
        </tr>
        <tr>
            <td>
                <b>cache_path</b>
            </td>
            <td>
                Path to extract the file or directory cached. For instance, if you used <a href="https://github.com/alephao/bitrise-step-s3-cache-push">S3 Cache Push</a> with the path <tt>./Carthage</tt> then this value should be <tt>./</tt>
            </td>
        </tr>
    </tbody>
</table>

#### Cache Key

The cache key.
