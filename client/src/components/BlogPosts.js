import React, { Component } from 'react';
import Axios from 'axios';
import _ from 'lodash';


const BlogPost = (props) => {
    return (
        <div>
            <h2>{ props.date }</h2>
            <p>{ props.post }</p>
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
        const data = await Axios.get('http://localhost:8000/api/blogposts');
        return await data;
    }

    componentDidMount() {
        this.getData().then(payload => {
            this.setState({
                posts: payload.data.posts
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
