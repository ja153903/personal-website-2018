import React, { Component } from 'react';

// have a class for now in case we want to put in state

class BlogEntryForm extends Component {

    constructor(props) {
        super(props);
        this.state = {
            blogPost: ''
        }
    }

    handleChange = (e) => {
        this.setState({
            blogPost: e.target.value
        });
    }

    handleSubmit = (e) => {
        e.preventDefault();
        console.log(this.state.blogPost);
    }

    render() {
        return (
            <form onSubmit={this.handleSubmit}>
                <label>Write Below</label>
                <div>
                    <input type="text" name="blog_post" onChange={this.handleChange} />
                </div>
                <button>Show the world</button>
            </form>
        )
    }
}

export default BlogEntryForm;
