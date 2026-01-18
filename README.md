# Getting started

## Variables
The project expects the following environment variables are set:
```bash
export COMINGSOON_PROTOCOL='http'
export COMINGSOON_HOST='localhost'
export COMINGSOON_PORT='8000'
export COMINGSOON_FOR_PROJECT='ComingSoon'
```

And this variable for building the docker image. You can use a private repository or docker's url:
```bash
export IMAGE_REGISTRY_BASE_URL=registry.hub.docker.com
```

## Running the API
Start here:
```bash
npm i
```

### For production
If you are building for production, run this command, otherwise you may skip it:
```bash
npm run build
```

To run the production build, run this command:
```bash
npm run serve
```

### For development
For development purposes we have two approaches, one for running with NPM and another for using Docker.

The NPM approach:
```bash
npm run dev
```

The Docker approach:
```bash
npm run docker
```
