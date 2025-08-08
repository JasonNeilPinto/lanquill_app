import React from "react";
import { Link } from "react-router-dom";
import { blogContent } from "../../data";

const Blogtest = () => {
  return (
    <>
      <section className="dg-blog-section bg-design-agency-about ptb-80">
        <div className="container">
          <div className="row justify-content-center">
            <div className="col-xl-5">
              <div className="section-title text-center mb-5">
                <span className="text-primary fw-bold">OUR BLOGS</span>
                <h2 className="heading-dg-color mb-0 mt-2 clr-text">
                  Our latest news
                </h2>
              </div>
            </div>
          </div>
          <div className="row g-4 justify-content-center">
            {blogContent.map((blog, index) => (
              <div className="col-xl-4" key={index}>
                <article className="dg-blog-card bg-white rounded-3">
                  <div className="thumbnail overflow-hidden rounded-3 mb-30">
                    <Link to={blog.blogPath}>
                      <img
                        src={blog.image}
                        alt="thumbnail"
                        className="img-fluid"
                      />
                    </Link>
                  </div>
                  <Link
                    to={blog.blogPath}
                    className="d-inline-block dg-text-color badge dg-bg-color"
                  >
                    {blog.blogType}
                  </Link>
                  <Link href={blog.blogPath}>
                    <h5 className="mb-3 mt-3">{blog.title}</h5>
                  </Link>
                  <p className="mb-4">
                    {blog.subtitle.length > 100
                      ? blog.subtitle.substring(0, 100) + "..."
                      : blog.subtitle}
                  </p>
                  <a
                    href={blog.blogPath}
                    className="read-more-link text-decoration-none dg-blog-btn-text"
                  >
                    Explore More <i className="fas fa-arrow-right ms-2"></i>
                  </a>
                </article>
              </div>
            ))}
          </div>
        </div>
      </section>
    </>
  );
};

export default Blogtest;
