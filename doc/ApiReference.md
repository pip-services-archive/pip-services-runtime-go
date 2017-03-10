# Pip.Services Runtime for Go API

Welcome to the Pip.Services Runtime for Golang API docs page. These pages contain the reference materials 
for the version 0.0.

The documentation is organized by classes and interfaces. They are groupped into packages below:

## [Core](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/core.html)

Contains interfaces for microservice components and other key abstractions.

- [IComponent](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/core.html#IComponent) - microservice component interface with component lifecycle definition
- [IDiscovery](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/core.html#IDiscovery) - interface for service discovery components
- [IConfig](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/core.html#IConfig) - interface for configuration readers
- [ILog](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/core.html#ILog) - interface for microservice loggers
- [ICounters](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/core.html#ICounters) - interface for performance counters
- [ICache](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/core.html#ICache) - interface for transient cache
- [IDataAccess](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/core.html#IDataAccess) - interface for data access (persistence) components
- [IController](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/core.html#IController) - interface for business logic controller
- [IInterceptor](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/core.html#IInterceptor) - interface for interceptors that decorate controller and customize microservice behavior
- [IClient](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/core.html#IClient) - interface for clients connectors to other microservices
- [IService](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/core.html#IService) - interface for API service (endpoint) that microservice exposes to its clients
- [IAddon](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/core.html#IAddon) - interface for microservice extentions 
- [References](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/core.html#References) - microservice component references
- [DynamicMap](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/core.html#DynamicMap) - language-independent hashmap to store hierarchical dynamic data
- [AbstractComponent](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/core.html#AbstractComponent) - abstract class for all microservice components
- [LogLevel](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/core.html#LogLevel) - enumeration for logging levels
- [Counter](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/core.html#Counter) - performance counter
- [CounterType](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/core.html#CounterType) - enumeration for performance counter types
- [Timing](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/core.html#Timing) - elapsed time measurement handler
- [MicroserviceError](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/core.html#MicroserviceError) - base class for all microservice errors
- [IIdentifiable](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/core.html#IIdentifiable) - interface for data objects identified by string **id**s
- [FilterParams](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/core.html#FilterParams) - free-form query parameters
- [PageParams](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/core.html#PageParams) - query paging parameters
- [DataPage](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/core.html#DataPage) - envelop object to return paged query result
 
## Build

Microservice construction and lifecycle management

- [LifeCycleManager](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/build.html#LifeCycleManager) - component lifecycle manager that inits, opens and closes a set of components
- [Builder](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/build.html#Build) - microservice builder that creates components following configuration as a build recipe

## Run

Microservice execution wrappers for various deployment platforms

- [Microservice](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/run.html#Microservice) - microservice instance
- [ProcessRunner](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/run.html#ProcessRunner) - runs microservices as a system process. It if also used for windows services, linux deamons or in docker deployments

## Discovery

Discovery components implementations. They used to register service addresses and dynamically 
resolve them for clients. 

- [AbstractDiscovery](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/discovery.html#AbstractDiscovery) - abstract class for all discovery components

## Config

Configuration components implementations. They used to define types of microservice components and
initialize their execution settings.

- [AbstractConfig](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/config.html#AbstractConfig) - abstract class for all config components
- [DirectConfig](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/config.html#DirectConfig) - exposes static / directly set microservice configuration
- [JsonConfig](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/config.html#JsonConfig) - reads micconfiguration configuration from JSON file on disk 

## Log

Logging components implementation. They used to report about microservice errors or log transactions
performance by microservice.

- [AbstractLog](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/log.html#AbstractLog) - abstract class for all logging components
- [NullLog](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/log.html#NullLog) - NULL logger that doesn't do anything
- [ConsoleLog](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/log.html#ConsoleLog) - logger that outputs messages to standard output and error streams
- [CompositeLog](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/log.html#CompositeLog) - log aggregator that passes messages to multiple loggers at once
- [CachedLog](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/log.html#CachedLog) - abstract logger that caches messages for batch output
- [LogEntry](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/log.html#LogEntry) - individual log entry used by CachedLogger

## Counters

Performance counting components implementations. They used to monitor microservice performance by measuring
time intervals, counting calls, recording timestamps of key events or arbitrary values.

- [AbstractCounters](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/counters.html) - abstract class for all performance counters
- [NullCounters](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/counters.html#NullCounters) - NULL performance counter that doesn't do anything
- [LogCounters](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/counters.html#LogCounters) - performance counter that periodically dumps counters to log output

## Cache

Caching components implementations. They used to store frequently accessed data in transient cache 
to avoid roundtrips to persistent storage for better performance.

- [AbstractCache](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/cache.html#AbstractCache) - abstract class for all caching components
- [NullCache](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/cache.html#NullCache) - NULL cache that doesn't do anything
- [MemoryCache](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/cache.html#MemoryCache) - local in-memory cache to be used for testing
- [CacheEntry](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/cache.html#CacheEntry) - cache value entry used by MemoryCache
- [MemcachedCache](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/cache.html#MemcachedCache) - connector to memcached distributed caching service https://memcached.org

## Db

Data access components implementations. They provide access to persistent storage.

- [AbstractDataAccess](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/db.html#AbstractDataAccess) - abstract class for all data access components
- [MemoryDataAccess](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/db.html#MemoryDataAccess) - in-memory data access to be used for testing
- [FileDataAccess](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/db.html#FileDataAccess) - file data access that reads and writes data from/to JSON file on disk

## Logic

Business logic controllers implementations. They encapsulate microservice business logic

- [AbstractController](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/logic.html#AbstractController) - abstract class for microservice controllers

## Ints

Interceptor components implementations. They decorate business logic controllers to modify their behavior to implement custom extensions.

- [AbstractInterceptor](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/ints.html#AbstractInterceptor) - abstract class for all interceptors

## Deps

Client dependency components implementations. They provide user-friendly connectors to microservice API endpoints.

- [AbstractClient](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/deps.html#AbstractClient) - abstract class for all client dependencies
- [RestClient](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/deps.html#RestClient) - interoperable HTTP/REST client

## Api

Service components implementations. They expose API endpoints to accept calls from microservice clients.

- [AbstractService](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/api.html#AbstractService) - abstract class for all API services
- [RestService](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/api.html#RestService) - interoperable HTTP/REST service

## Addons

Addon components implementations. They represent microservice extensions that take no part in business transactions, 
but provide additional services like reporting microservice health state, collecting usage metrics or performing 
random shutdowns for relience testing.

- [AbstractAddon](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/addons.html#AbstractAddon) - abstract class for all addons   

## Util

Package containing cross-language abstractions and utility functions used across microservice implementations.

- [Converter](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/util.html#Converter) - value type converter to convert objects to strings, numbers, dates and maps
- [IdGenerator](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/util.html#IdGenerator) - generator of unique object ids
- [TagsProcessor](http://htmlpreview.github.io/?https://github.com/pip-services/pip-services-runtime-go/blob/master/doc/api/util.html#TagProcessor) - extracts from data objects and processes search tags
