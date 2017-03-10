docgen:
	godoc -html github.com/pip-services/pip-services-runtime-go > doc/api/core.html
	godoc -html github.com/pip-services/pip-services-runtime-go/addons > doc/api/addons.html
	godoc -html github.com/pip-services/pip-services-runtime-go/api > doc/api/api.html
	godoc -html github.com/pip-services/pip-services-runtime-go/build > doc/api/build.html
	godoc -html github.com/pip-services/pip-services-runtime-go/cache > doc/api/cache.html
	godoc -html github.com/pip-services/pip-services-runtime-go/config > doc/api/config.html
	godoc -html github.com/pip-services/pip-services-runtime-go/counters > doc/api/counters.html
	godoc -html github.com/pip-services/pip-services-runtime-go/db > doc/api/db.html
	godoc -html github.com/pip-services/pip-services-runtime-go/deps > doc/api/deps.html
	godoc -html github.com/pip-services/pip-services-runtime-go/discovery > doc/api/discovery.html
	godoc -html github.com/pip-services/pip-services-runtime-go/ints > doc/api/ints.html
	godoc -html github.com/pip-services/pip-services-runtime-go/log > doc/api/log.html
	godoc -html github.com/pip-services/pip-services-runtime-go/logic > doc/api/logic.html
	godoc -html github.com/pip-services/pip-services-runtime-go/run > doc/api/run.html
	godoc -html github.com/pip-services/pip-services-runtime-go/util > doc/api/util.html
