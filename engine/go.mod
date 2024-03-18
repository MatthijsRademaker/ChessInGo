module chess/engine

go 1.22.1

replace chess/domain => ../domain

require chess/domain v0.0.0-00010101000000-000000000000

replace chess/domain/pieces => ../domain/pieces

require chess/domain/pieces v0.0.0-00010101000000-000000000000

replace chess/bot => ../bot

require chess/bot v0.0.0-00010101000000-000000000000
