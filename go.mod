module postgre-sample

go 1.20

require (
	github.com/lib/pq v1.10.9
)

replace flo.com.tr/types => ./types
require flo.com.tr/types v0.0.0-00010101000000-000000000000
