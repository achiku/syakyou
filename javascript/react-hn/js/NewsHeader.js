var $ = require('jquery');
var React = require('react');

var NewsHeader = React.createClass({
  getNav: function() {
    var navLinks = [
      {name: 'new', url: 'newest'},
      {name: 'comments', url: 'newcomments'},
      {name: 'show', url: 'show'},
      {name: 'ask', url: 'ask'},
      {name: 'jobs', url: 'jobs'},
      {name: 'submit', url: 'submit'},
    ];

    return (
      <div className="newsHeader-nav">
      </div>
    )
  },
  getLogo: function() {
    return (
      <div className="newsHeader-logo">
        <a href="https://www.ycombinator.com"><img src="../img/y18.gif"/></a>
      </div>
    );
  },
  getTitle: function() {
    return (
      <div className="newsHeader-title">
        <a className="newsHeader-textLink" href="https://news.ycombinator.com">Hacker News</a>
      </div>
    );
  },
  render: function() {
    return (
      <div className="newsHeader">
        {this.getLogo()}
        {this.getTitle()}
      </div>
    );
  },
});

module.exports = NewsHeader;
