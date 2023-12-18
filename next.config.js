const nextConfig = {
  async redirects() {
    // When running Next.js via Node.js (e.g. `dev` mode), proxy API requests
    // to the Go server.
    console.log("sdsf");
    return [
      {
        source: "/api",
        destination: "http://localhost:8080/api",
        permanent: true,
      },
    ];
  },
  // future: {
  //   webpack5: true,
  // },
  trailingSlash: true,
};

module.exports = nextConfig;