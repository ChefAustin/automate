import {
  Component, Input, OnChanges, ViewChild, ElementRef
} from '@angular/core';
import * as d3 from 'd3';


@Component({
  selector: 'app-simple-line-graph',
  templateUrl: './simple-line-graph.component.html',
  styleUrls: ['./simple-line-graph.component.scss']
})

export class SimpleLineGraphComponent implements OnChanges {

  constructor(
    private chart: ElementRef
  ) {}

  @Input() data: any = [];
  @Input() width = 900;
  @Input() height = 200;

  @ViewChild('svg', {static: true}) svg: ElementRef;

  public heightMargin = 40;
  public theHighlight;
  public locked: number = null;
  public currentHighlightnum: number;

  ////////   X AXIS ITEMS   ////////
  // maps all of our x data points
  get xData() {
    return this.data.map(d => d.daysAgo);
  }
  // determines how wide the graph should be to hold our data
  // in its respective area;
  get rangeX() {
    const min = 50;
    const max = this.width - 50;
    return [max, min];  // we want to plot our data backwards, so we reverse [min, max]
  }
  // determines the min and max values of the x axis
  get domainX() {
    const min = Math.min(...this.xData);
    const max = Math.max(...this.xData);
    return [min, max];
  }
  // determines each of our X axis points using the height and width of the chart
  get xScale() {
    return d3.scaleTime()
      .range(this.rangeX)
      .domain(this.domainX);
  }


  ////////   Y AXIS ITEMS   ////////

  // maps all of our Y data points
  get yData() {
    return this.data.map(d => d.percentage);
  }

  get rangeY() {
    const min = 10;
    const max = this.height - this.heightMargin;
    return [max, min];
  }

  get domainY() {
    const min = 0;
    const max = 100; // since this based on a percentage do we want 0 to 100?
    return [min, max];
  }
  // determines each of our Y axis points using the height and width of the chart
  get yScale() {
    return d3.scaleLinear()
             .range(this.rangeY)
             .domain(this.domainY);
  }

  //////// SELECTIONS ////////
  get containerSelection() {
    return d3.select('app-simple-line-graph');
  }

  get axisYSelection() {
    return this.svgSelection.select('.y-axis');
  }

  get svgSelection() {
    return d3.select(this.svg.nativeElement);
  }

  get theToolTip() {
    return d3.select('chef-tooltip');
  }

  get theRectHighlight() {
    return d3.select('.rect-highlight');
  }

  // returns a function that when passed our data, will return an svg path
  get createPath() {
    return d3.line()
              .x(d => this.xScale( d.daysAgo) )
              .y(d => this.yScale( d.percentage) );
  }

  get viewBox() {
    return `0 0 ${this.width} ${this.height}`;
  }

    ////////////////// RENDER FUNCTIONS ////////////////////
  renderChart() {
    this.AllUnlockDeactivate();
    this.resizeChart();
    this.renderGrid();
    this.renderLine();
    this.renderPoints();
    this.renderTooltips();
    this.renderRings();
    this.renderLabelButtons();
    this.relock();
  }

  renderLine(): void {
    // create the line using path function
    const line = this.createPath(this.data);

    const theLine = this.svgSelection.selectAll('.line').data([this.data],d => d.daysAgo);
    theLine.exit().remove();
    theLine.enter().append('path').attr('class', 'line').merge(theLine)
    .transition().duration(1000)
    .attr('d', line);
  }

  renderPoints(): void {
    const points = this.svgSelection.selectAll('circle.point')
      .data(this.data, d => d.daysAgo);
    points.exit().remove();
    points.enter().append('circle')
        .attr('class', (_d,i) => `point elem-${i}`)
        .merge(points)
        .transition().duration(1000)
        .attr('percent', ( d => d.percentage ) ) // must add this data AFTER the merge
        .attr('cx', d => this.xScale(d.daysAgo))
        .attr('cy', d => this.yScale(d.percentage))
        .attr('r', 4);
  }

  renderRings(): void {
    const rings = this.svgSelection.selectAll('circle.ring')
      .data(this.data, d => d.daysAgo);
    rings.exit().remove();
    rings.enter().append('circle')
      .attr('class', (_d, i) => `ring elem-${i}`)
      .merge(rings)
      .attr('cx', d => this.xScale(d.daysAgo))
      .attr('cy', d => this.yScale(d.percentage))
      .attr('r', 10);
  }

  renderTooltips() {
    const tooltips = this.containerSelection.selectAll('div.graph-tooltip')
      .data(this.data, d => d.daysAgo);
    tooltips.exit().remove();
    tooltips.enter().append('div')
      .attr('class', (_d, i) => `graph-tooltip elem-${i}`)
      .merge(tooltips)
      .text(d => `Checked in ${d.percentage}%`)
      .style('left', d => `${this.xScale(d.daysAgo) - 50}px`) // magic number...for now
      .style('top', d => `${this.yScale(d.percentage) + 20}px`);
  }

  renderLabelButtons() {
    const labels = this.containerSelection.selectAll('.graph-button')
    .data(this.data, d => d.daysAgo);
    labels.exit().remove();
    labels.enter().append('button')
      .call(parent => {
        parent.append('div')
          .attr('class', (_d, i) => `inner elem-${i}`);
      })
      .attr('class', (_d, i) => `graph-button elem-${i}`)
      .merge(labels)
      .transition().duration(1000)
      .style('top', `calc(100% - ${this.heightMargin}px)`)
      .style('left', d => `${this.xScale(d.daysAgo) - this.heightMargin}px`) // will need adjustment
      .call(parent => {
        parent.select('.inner')
          .text(p => this.formatLabels(p.daysAgo));
      });

      // add all listeners
    this.containerSelection.selectAll('.graph-button')
      .classed('turnt', () => this.xData.length > 7)

      .on('mouseenter', () => {
        this.handleHover(d3.event);
      })
      .on('mouseout', () => {
        this.AllDeactivate();
      })
      // focus styles
      .on('focus', () => {
        this.handleHover(d3.event);
      })
      .on('focusout', () => {
        this.AllDeactivate();
      })
      .on('click', () => {
        this.handleClick(d3.event);
      });
  }

  formatLabels(daysAgo): string {
    switch (daysAgo) {
      case 0:
        return '24 hrs ago';
        break;
      default:
        return `${daysAgo + 1} days ago`;
    }
  }

  renderGrid() {
    // create the X axis grid lines
    const xGrid = d3.axisTop()
      .ticks(this.data.length)
      .tickFormat('')
      .tickSize(this.height - ( this.heightMargin * ( 3 / 2 ) ) + 10) // magic numbers?
      .tickSizeOuter(0)
      .scale(this.xScale);
    // Render the X axis and X ticks
    const grid = this.svgSelection.selectAll('.grid').data([this.data]);
    grid.exit().remove();
    grid.enter().append('g').attr('class', 'grid')
      // this line will need to be updated to flexible
      .attr('transform', `translate(0, ${this.height - this.heightMargin})`)
      .merge(grid).transition().duration(1000)
      .call(xGrid);

    // create the Y axis
    const yAxis = d3.axisRight(this.yScale).tickFormat(d => d + '%').ticks(1);
    // render the Y axis
    const y = this.svgSelection.selectAll('.y-axis').data([this.data]);
    y.exit().remove();
    y.enter().append('g').attr('class', 'y-axis').merge(y)
      .transition().duration(1000)
      .call(yAxis);

    const xAxis = d3.axisBottom().ticks(this.data.length)
      .tickSizeInner(10).tickSizeOuter(0).tickFormat('')
      .scale(this.xScale);

    const x = this.svgSelection.selectAll('.x-axis').data([this.data]);
    x.exit().remove();
    x.enter().append('g').attr('class', 'x-axis')
      .attr('transform', `translate(0, ${this.height - this.heightMargin})`)
      .merge(x).transition().duration(1000)
      .call(xAxis);

    // remove zero from bottom of chart on x axis
    this.svgSelection.selectAll('.tick')
      .filter(tick => tick === 0)
      .remove();
  }

  handleHover(d3Event): void {
    const num = this.getHoveredElement(d3Event);
    d3.selectAll(`.elem-${num}`).classed('active', true);
  }

  handleClick(d3Event): void {
    const num = this.getHoveredElement(d3Event);
    const isAlreadyLocked = d3.selectAll(`.elem-${num}`).classed('lock');
    if ( isAlreadyLocked ) {
      d3.selectAll(`.elem-${num}`).classed('lock', false);
      this.locked = null;
    } else {
      d3.selectAll('.lock').classed('lock', false);
      d3.selectAll(`.elem-${num}`).classed('lock', true);
      this.locked = num;
    }
  }

  getHoveredElement(d3Event): number {
    const classes = d3.select(d3Event.target).attr('class');
    const match = classes.match(/elem-([0-9]{1,2})/g)[0];
    const num = match.split('-')[1];
    return num;
  }

  relock(): void {
    console.log(this.locked);
    if (this.locked) {
      console.log('isLocked');
      d3.selectAll(`.elem-${this.locked}`).classed('lock', true);
    }
  }

  AllDeactivate(): void {
    d3.selectAll('.active').classed('active', false); // deactivate any active items on page
  }

  allUnlock(): void {
    d3.selectAll('.lock').classed('lock', false); // unlock any locked items on page
  }

  AllUnlockDeactivate(): void {
    this.allUnlock();
    this.AllDeactivate(); // deactivate any active items on page
  }


  ngOnChanges() {
    this.renderChart();
  }

  onResize(): void {
    this.renderChart();
  }

  resizeChart(): void {
    this.width = this.chart.nativeElement.getBoundingClientRect().width;
  }

}

