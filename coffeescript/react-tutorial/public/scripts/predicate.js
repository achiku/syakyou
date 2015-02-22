var Counter, React, jade, _,
  __bind = function(fn, me){ return function(){ return fn.apply(me, arguments); }; },
  __hasProp = {}.hasOwnProperty,
  __extends = function(child, parent) { for (var key in parent) { if (__hasProp.call(parent, key)) child[key] = parent[key]; } function ctor() { this.constructor = child; } ctor.prototype = parent.prototype; child.prototype = new ctor(); child.__super__ = parent.prototype; return child; };

React = require('react');

jade = require('react-jade');

_ = require('lodash');

Counter = (function(_super) {
  __extends(Counter, _super);

  function Counter() {
    this.render = __bind(this.render, this);
    this.tick = __bind(this.tick, this);
    this.state = {
      count: 0
    };
  }

  Counter.prototype.tick = function() {
    return this.setState({
      count: this.state.count + 1
    });
  };

  Counter.prototype.render = function() {
    return jade.compile("#counter\n  span Count :\n  button(onClick=tick)= count")(_.assign({}, this, this.props, this.state));
  };

  return Counter;

})(React.Component);

React.render(React.createFactory(Counter)(), document.getElementById('container'));
