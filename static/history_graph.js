const response = fetch("/history")
  .then((response) => response.text())
  .then((body) => JSON.parse(body))
  .catch((err) => console.log(err));

response.then((data) => {
  console.log("history", data);

  const events = data.events;
  const players = data.players;
  const margin = 40;

  const I = d3.range(Object.values(events)[0].length);

  const xDomain = [0, Object.values(events)[0].length - 1];
  const yDomain = d3.extent(
    Object.values(events)
      .flat()
      .map((v) => v.rating)
  );

  const width = 400;
  const height = 600;
  const xRange = [margin, width - margin];
  const yRange = [margin, height - margin];
  const xScale = d3.scaleLinear(xDomain, xRange);
  const yScale = d3.scaleLinear(yDomain, yRange);
  const yAxis = d3.axisLeft(yScale).ticks(height / 40);

  const svg = d3
    .select("#canvas")
    .append("svg")
    .attr("viewBox", [0, 0, width, height])
    .attr("style", "max-width: 100%; height: auto; height: intrinsic;");

  svg
    .selectAll("image")
    .data(Object.keys(events))
    .enter()
    .append("image")
    .attr("x", 5 + width - margin)
    .attr("y", (d) => yScale(events[d][events[d].length - 1].rating) - 15)
    .attr("width", 30)
    .attr("height", 30)
    .attr(
      "href",
      (d) => `/static/characters/${players.find((p) => p.id == d).icon}.png`
    );

  // Draw yAxis
  svg
    .append("g")
    .attr("transform", `translate(${margin},0)`)
    .call(yAxis)
    .call((g) => g.select(".domain").remove())
    .call((g) =>
      g
        .selectAll(".tick line")
        .clone()
        .attr("x2", width)
        .attr("stroke-opacity", 0.1)
    );

  // Draw players history
  for (const index in events) {
    const Y = d3.map(events[index], (d) => d.rating);

    const line = d3
      .line()
      .x((i) => xScale(i))
      .y((i) => yScale(Y[i]))
      .curve(d3.curveMonotoneX);

    svg
      .append("path")
      .attr("fill", "none")
      .attr("stroke", d3.interpolateRainbow(index / players.length))
      .attr("stroke-width", 3)
      .attr("stroke-linecap", "round")
      .attr("stroke-linejoin", "round")
      .attr("d", line(I));
  }

//   d3.select("#canvas")
//     .selectAll("p")
//     .data(players)
//     .enter()
//     .append("p")
//     .style("color", (d) => d3.interpolateRainbow(d.id / players.length))
//     .text((d) => d.name);
});
