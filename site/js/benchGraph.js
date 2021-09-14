var BenchGraph = (function () {

	function Render() {
		for (var i = 0; i < this.data.length; i++) {
			var bench = this.data[i];
			var graphContainer = $('.graph.' + bench.Program);
			if (graphContainer.length == 0) {
				console.warn('no container found for ' + bench.Program + ' graph');
				continue
			}
			graphContainer.find('.program span').text(bench.Program.replace('-', ' '));
			var maxTime = 0;
			var minTime = 9999999999999;
			for (const interpreter in bench.Results) {
				var res = bench.Results[interpreter];
				if (res.Time > maxTime) {
					maxTime = res.Time;
				}
				if (res.Time < minTime) {
					minTime = res.Time;
				}
			}
			var times = [];
			var allocs = [];
			var timeFormat = 's';
			if ( minTime / 1000000000 < 1 ) {
				timeFormat = 'ms';
			}
			if ( minTime / 1000000 < 1 ) {
				timeFormat = 'μs'
			}
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

				var timeLabel;
				if ( timeFormat == 's' ) {
					timeLabel = (Math.round((res.Time / 1000000000 + Number.EPSILON) * 100) / 100) + 's';
				} else if ( timeFormat == 'ms' ) {
					timeLabel = Math.round(res.Time / 1000000) + 'ms';
				} else {
					timeLabel = Math.round(res.Time / 1000) + 'μs';
				}
				var timeWidth = res.Time * 90 / maxTime;
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

			for (var j = 0; j < times.length; j++) {
				var langContainer = graphContainer.children(' .' + times[j].interpreter);
				langContainer.show();
				graphContainer.append(langContainer);
			}
			for (var j = 0; j < allocs.length; j++) {
				var langContainer = graphContainer.find('.allocations .data .' + allocs[j].interpreter);
				langContainer.css({ display : 'table-row' });
				graphContainer.find('.allocations .data').append(langContainer);
			}
			graphContainer.append(graphContainer.find('.allocations'));
		}

	}

	// Constructor
	function benchGraph(benchmarksData) {
		this.data = benchmarksData;
		this.Render = Render;
	}

	return benchGraph
})();