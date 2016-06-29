
```
gcloud config list
gcloud config set project <project_id>

gsutil ls

# creating a bucket
gsutil mb gs://job-data-development

# upload an object into the bucket
gsutil cp <source> gs://job-data-development

# downloading an object from the bucket
gsutil cp gs://job-data-development/storage.md /tmp

# copying an object to a folder in the bucket
gsutil cp gs://job-data-development/storage.md gs://job-data-development/some_folder/

# list contents of a bucket or folder
gsutil ls gs://job-data-development

# to list details use -l switch
gsutil ls -l gs://job-data-development/storage.md

# make your object publicly accessible
gsutil acl ch -u AllUsers:R gs://job-data-development/storage.md

# to remove the permission
gsutil acl ch -d AllUsers:R gs://job-data-development/storage.md

# give someone access to the bucket
gsutil acl ch -u <email>:W gs://job-data-development

# to remove the permission
gsutil acl ch -d <email> gs://job-data-development

# removing objects
gsutil rm gs://job-data-development/storage.md

# to remove a bucket
gsutil rm -r gs://job-data-development/storage.md
```

Once a bucket has been created, storage class and location can only be
changed by deleting and re-creating the bucket.

Objects are the individual pieces of data that we store in Google Cloud
Storage. A single object can be upto 5TB in size. Objects have two
components: object data and object metadata.
 * object data is the file
 * object metadata is a collection of name-value pairs that describe
   various object qualities.

By using slashes in an object name, we can make objects appear as though
they're stored in a hierarchical strucure. But Google Cloud Storage sees
the objects as independent objects with no hierarchical relationship.

* objects are immutable.
* incremental changes to the object is not possible.
* object overwrite can be done. It is done atomically, until the new
  upload completes the old version of the object will be served to
  readers.

There is only one Google Cloud Storage namespace, which means every
bucket must have a unique name across the entire google cloud storage
namespace. Object names must be unique only within a given bucket.

### Access Control


### Google Cloud Storage Authentication

#### OAuth 2.0
  GCS uses OAuth 2.0 for API authentication and authorization.
  Authentication is the process of determining the identity of a client.
  The details of authentication vary depending on how we access GCS, but
  fall into 2 general types:
  * A server-centric flow allows an application to directly hold the
    credentials of a service account to complete authentication. Use
    this flow if application works with its own data rather than user
    data. Google Cloud Platform projects have default service accounts
    we can use or we can create new ones.

  * A user-centric flow allows an application to obtain credentials from
    an end user. The user signs in to complete authentication. Use this
    flow if application needs to access user data.

  OAuth uses scopes to determine if an authenticated identity is
  authorized.

#### gsutil authentication
  With gsutil installed from the cloud SDK, we can authenticate with
  user or service account credentials.
  ` gcloud auth activate-service-account`
  `gcloud auth` uses the cloud-platform scope when getting an access
  token (view and manage data across all gcp services).


#### Client library authentication
  Client libraries can use application default credentials to easily
  authenticate with Google APIs and send requests to those APIs.

  Application Default Credentials are part of the client libraries we
  can use to access GCS. Default credentials identify the application
  with either a user credential or a default service account.

#### Service Account Credentials
  Service accounts are special accounts that represent software rather
  than people. They are the most common way applications authenticate
  with GCS.

  When we use a service account to authenticate the application, we do
  not need a user to authenticate to get an access token. Instead, we
  obtain a private key from google cloud platform console, we we then
  use to send a signed request for an access token. We can then use the
  access token like we normally would.

- Create a service account from cloud console in JSON.
