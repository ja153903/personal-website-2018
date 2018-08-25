import React, { Component } from 'react';

// have a class for now in case we want to put in state

class BlogEntryForm extends Component {

    onSubmit = (e) => {
        console.log(e.target.value);
    }

    render() {
        return (
            <form onSubmit={this.onSubmit}>
                <label>Write Below</label>
                <input type="text" />
                <input type="submit" value="Show the world"/>
            </form>
        )
    }
}
