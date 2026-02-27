# oph.comingsoon.api

A coming soon API container project used as a placeholder in the project: [oph.tf.env](https://github.com/onlineproducthouse/oph.tf.env)

## Installation

```bash
# clone repository
mkdir oph.comingsoon.api
cd oph.comingsoon.api
git clone https://github.com/onlineproducthouse/oph.comingsoon.api.git .

# install dependencies
npm i

# build project
npm run build
```

## Usage

```bash
# set to either: local, test, qa, prod
export ENVIRONMENT_NAME=local

# set environment variables for the postgres database instance
export COMINGSOON_PROTOCOL='http'
export COMINGSOON_HOST='localhost'
export COMINGSOON_PORT='8000'
export COMINGSOON_FOR_PROJECT='ComingSoon'

# run the API
npm run serve

# run the API for development purposes
npm run serve
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
