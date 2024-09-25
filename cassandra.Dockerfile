# Use the official Cassandra image as a parent image
FROM cassandra:latest

# Set the working directory in the container
WORKDIR /var/lib/cassandra

# Copy any configuration files if needed
# COPY ./my-custom-config.yaml /etc/cassandra/cassandra.yaml

# Expose the default Cassandra ports
EXPOSE 7000 7001 7199 9042 9160

# Set the entry point to run Cassandra
CMD ["cassandra", "-f"]

# docker build -t my-cassandra -f cassandra.Dockerfile .
# docker run -d --name cassandra-node -p 9042:9042 my-cassandra
