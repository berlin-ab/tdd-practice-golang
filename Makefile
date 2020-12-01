#
# TDD Practice
#

test-rock-paper-scissors:
	ginkgo -r ./internal/rockpaperscissors

watch-rock-paper-scissors:
	ginkgo watch -r ./internal/rockpaperscissors
