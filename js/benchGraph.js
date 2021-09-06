var BenchGraph = (function () {

	function Render() {
		for (var i = 0; i < this.data.length; i++) {
			var bench = this.data[i];
			var graphContainer = $('.graph.' + bench.Program);
			if (graphContainer.length == 0) {
				console.warn('no container found for ' + bench.Program + ' graph');
				continue
			}
			graphContainer.find('.program span').text(bench.Program);
			var maxTime = 0;
			for (const interpreter in bench.Results) {
				var res = bench.Results[interpreter];
				if (res.Time > maxTime) {
					maxTime = res.Time
				}
			}
			var times = [];
			var allocs = [];
			for (const interpreter in bench.Results) {
				var res = bench.Results[interpreter];
				times.push({
					interpreter: interpreter,
					value: res.Time
				});
				allocs.push({
					interpreter: interpreter,
					value: res.Allocs
				});

				var timeLabel = (Math.round((res.Time / 1000000000 + Number.EPSILON) * 100) / 100) + 's';
				var timeWidth = res.Time * 95 / maxTime;
				graphContainer.find('.' + interpreter + ' .bar').css({
					width: timeWidth + '%'
				});
				graphContainer.find('.' + interpreter + ' span.time').text(timeLabel)

				var allocsLabel = graphContainer.find('.allocations .' + interpreter + ' .allocs');
				var allocsText = '';
				if (res.Allocs > 1000000) {
					allocsText = '~ ' + Math.round(res.Allocs / 1000000) + " million";
				} else {
					allocsText = res.Allocs;
				}
				allocsLabel.text(allocsText);
			}

			function sortFunc(a, b) {
				return a.value - b.value;
			}
			times.sort(sortFunc);
			allocs.sort(sortFunc);

			for (var i = 0; i < times.length; i++) {
				graphContainer.append(graphContainer.children(' .' + times[i].interpreter));
			}
			for (var i = 0; i < allocs.length; i++) {
				graphContainer.find('.allocations .data').append(graphContainer.find('.allocations .data .' + allocs[i].interpreter));
			}
			// $('.benchmark .graph').append($('.benchmark .graph .allocations'));
		}

	}

	// Constructor
	function benchGraph(benchmarksData) {
		this.data = benchmarksData;
		this.Render = Render;
	}

	return benchGraph
})();