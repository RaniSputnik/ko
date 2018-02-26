FROM ruby:latest

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