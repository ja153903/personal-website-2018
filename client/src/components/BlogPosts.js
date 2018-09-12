import React, { Component } from 'react';
import Axios from 'axios';
import _ from 'lodash';
import { Markdown } from 'react-showdown';


// Assume that the post we get will be in markdown
const BlogPost = (props) => {
    return (
        <div>
            <h1>{ props.date }</h1>
            <Markdown markup={props.post} />
        </div>
    )
}

class BlogPosts extends Component {
    constructor(props) {
        super(props);
        this.state = {
            posts: {}
        }
    }

    getData = async () => {
        const data = await Axios.get('http://localhost:5000/api/blogpost');
        return await data;
    }

    componentDidMount() {
        this.getData().then(payload => {
            this.setState({
                posts: payload.data
            });
        }).catch(err => {
            console.log(err);
        });
    }

    renderPosts = () => {
        let result = [];
        _.forEach(this.state.posts, (value, key) => {
            result.push((
                <li key={key}>
                    <BlogPost
                        date={key}
                        post={value} />
                </li>
            ))
        })
        return result;
    }

    render() {
        return(
            <div>
                { this.renderPosts() }
            </div>
        )
    }
}

export default BlogPosts;
