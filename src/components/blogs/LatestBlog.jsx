import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import SectionTitle from "../common/SectionTitle";

const initialBlogData = [
  {
    id: 1,
    title: "A Beginner’s Guide to IT Outsourcing in Bangalore",
    description:
      "Bangalore, often referred to as the Silicon Valley of India, is a global hub for technology and innovation. With its thriving IT ecosystem, the city has become a hotspot for businesses looking to outsource their IT needs.",
    image: "/img/blog/beginnerit.png",
    category: "Blog",
    categoryClass: "bg-primary-soft",
    author: {
      name: "Jane Martin",
      avatar: "/img/testimonial/6.jpg",
      date: "April 24, 2021",
    },
    link: "/blog-single",
  },
  {
    id: 2,
    title: "Talent Acquisition Trends 2025 fit",
    description:
      "As we approach 2025, talent acquisition is poised to undergo transformative changes driven by technology, evolving workforce expectations, and global economic shifts.",
    image: "/img/blog/Talent.png",
    category: "Blog",
    categoryClass: "bg-primary-soft",
    author: {
      name: "Veronica P. Byrd",
      avatar: "/img/testimonial/1.jpg",
      date: "April 24, 2021",
    },
    link: "/blog-second",
  },
  {
    id: 3,
    title: "Shaping India's Future: 12 Big Ideas for 2025",
    description:
      "As India continues its upward trajectory, it's imperative for business leaders to stay ahead of the curve and anticipate the trends that will shape the country's future.",
    image: "/img/blog/big-ideas.png",
    category: "Blog",
    categoryClass: "bg-primary-soft",
    author: {
      name: "Martin Gilbert",
      avatar: "/img/testimonial/3.jpg",
      date: "April 24, 2021",
    },
    link: "/blog-ShapingIndia",
  },
];

const LatestBlog = () => {
  const [blogs] = useState(initialBlogData);

  useEffect(() => {}, []);

  return (
    <section className="related-blog-list ptb-80 bg-light">
      <div className="container">
        <div className="row align-items-center justify-content-between">
          <div className="col-lg-4 col-md-12">
            <SectionTitle
              subtitle="Blog"
              title="Check our Latest News and Update"
            />
          </div>
          <div className="col-lg-7 col-md-12">
            <div className="text-start text-lg-end mb-4 mb-lg-0 mb-xl-0">
              <Link to="/blogs" className="btn btn-primary">
                View All Article
              </Link>
            </div>
          </div>
        </div>
        <div className="row">
          {blogs.map((blog) => (
            <div key={blog.id} className="col-lg-4 col-md-6 d-flex">
              <div className="single-article rounded-custom mb-4 mb-lg-0">
                <Link to={blog.link} className="article-img">
                  <img src={blog.image} alt="article" className="img-fluid" />
                </Link>
                <div className="article-content p-4">
                  <div className="article-category mb-4 d-block">
                    <Link
                      to="#"
                      className={`d-inline-block text-dark badge ${blog.categoryClass}`}
                    >
                      {blog.category}
                    </Link>
                  </div>
                  <Link to={blog.link}>
                    <h2 className="h5 article-title limit-2-line-text">
                      {blog.title}
                    </h2>
                  </Link>
                  <p className="limit-2-line-text">{blog.description}</p>
                  <Link to="#" className="mt-auto">
                    <div className="d-flex align-items-center pt-4">
                      <div className="avatar">
                        <img
                          src={blog.author.avatar}
                          alt="avatar"
                          width="40"
                          className="img-fluid rounded-circle me-3"
                        />
                      </div>
                      <div className="avatar-info">
                        <h6 className="mb-0 avatar-name">{blog.author.name}</h6>
                        <span className="small fw-medium text-muted">
                          {blog.author.date}
                        </span>
                      </div>
                    </div>
                  </Link>
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </section>
  );
};

export default LatestBlog;
