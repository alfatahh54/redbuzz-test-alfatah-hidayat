const nextConfig = {
  async redirects() {
    // When running Next.js via Node.js (e.g. `dev` mode), proxy API requests
    // to the Go server.
    console.log("sdsf");
    return [
      {
        source: "/api/:path*",
        destination: "http://localhost:8080/api/:path*",
        permanent: true,
      },
    ];
  },
  trailingSlash: true,
};

module.exports = nextConfig;