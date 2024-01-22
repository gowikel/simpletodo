# Simpletodo

I created this app to explore the concept of Domain Driven Design using golang.

### Building and running

The application uses docker compose for an easy local deployment. To build, simply
run:

```bash
docker compose up --build
```

And it will be available at http://localhost:8080.

If crosscompilation is needed (e.g a cloud deployment where you are using a Mac M1 but the provider uses a amd64), you can target it with the `--platform` flag:

```bash
docker build --platform=linux/amd64 -t simpletodo .
```