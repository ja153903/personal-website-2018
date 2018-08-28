import React, { Component } from 'react';
import Axios from 'axios';

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
        // Keep the local env until we get the API hosted outside
        Axios.post('http://localhost:8000/api/blogposts', {
            'blog_post': this.state.blogPost,
            'date': (new Date()).toString(),
        }).then(res => {
            console.log(res);
        }).catch(err => {
            console.log(err);
        });
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
