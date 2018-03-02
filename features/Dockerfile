FROM ruby:latest

# Install mysql client
RUN apt-get update \
  && apt-get install -y mysql-server mysql-client default-libmysqlclient-dev --no-install-recommends \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Create src/features directory
RUN mkdir /usr/src/features
WORKDIR /usr/src/

# Install dependencies
ADD Gemfile Gemfile
ADD Gemfile.lock Gemfile.lock
RUN bundle install

# Add feature tests
ADD . /usr/src/features

# Run cucumber
CMD ["cucumber"]