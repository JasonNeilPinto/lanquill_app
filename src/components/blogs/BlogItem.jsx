import React from "react";
import { Link } from "react-router-dom";

const BlogItems = () => {
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
            <div className="col-xl-4">
              <article className="dg-blog-card bg-white rounded-3">
                <div className="thumbnail overflow-hidden rounded-3 mb-30">
                  <Link href="/blog-single">
                    <img
                      src="/img/blog/blog-1.jpg"
                      alt="thumbnail"
                      className="img-fluid"
                    />
                  </Link>
                </div>
                <Link
                  href="/blog"
                  className="d-inline-block dg-text-color badge dg-bg-color"
                >
                  Case Study
                </Link>
                <Link href="/blog-single">
                  <h5 className="mb-3 mt-3">
                    Continually reinterme diate seamless manufactured products.
                  </h5>
                </Link>
                <p className="mb-4">
                  Globally evolve user-centric channels rather than reliable
                  solutions. Globally strategize.
                </p>
                <a
                  href="#"
                  className="read-more-link text-decoration-none dg-blog-btn-text"
                >
                  Explore More <i className="fas fa-arrow-right ms-2"></i>
                </a>
              </article>
            </div>
            <div className="col-xl-4">
              <article className="dg-blog-card bg-white rounded-3">
                <div className="thumbnail overflow-hidden rounded-3 mb-30">
                  <Link href="/blog-single">
                    <img
                      src="/img/blog/blog-2.jpg"
                      alt="thumbnail"
                      className="img-fluid"
                    />
                  </Link>
                </div>
                <Link
                  href="blog"
                  className="d-inline-block dg-text-color badge dg-bg-color"
                >
                  Digital Marketing
                </Link>
                <Link href="/blog-single">
                  <h5 className="mb-3 mt-3">
                    5 SEO Factors to Consider for your Website to rank Better
                  </h5>
                </Link>
                <p className="mb-4">
                  Globally evolve user-centric channels rather than reliable
                  solutions. Globally strategize.
                </p>
                <a
                  href="#"
                  className="read-more-link text-decoration-none dg-blog-btn-text"
                >
                  Explore More <i className="fas fa-arrow-right ms-2"></i>
                </a>
              </article>
            </div>
            <div className="col-xl-4">
              <article className="dg-blog-card bg-white rounded-3">
                <div className="thumbnail overflow-hidden rounded-3 mb-30">
                  <Link href="/blog-single">
                    <img
                      src="/img/blog/blog-3.jpg"
                      alt="thumbnail"
                      className="img-fluid"
                    />
                  </Link>
                </div>
                <Link
                  href="/blog"
                  className="d-inline-block dg-text-color badge dg-bg-color"
                >
                  Business Planning
                </Link>
                <Link href="/blog-single">
                  <h5 className="mb-3 mt-3">
                    Quickly generate low risk high yield deliverables.
                  </h5>
                </Link>
                <p className="mb-4">
                  Globally evolve user-centric channels rather than reliable
                  solutions. Globally strategize.
                </p>
                <a
                  href="#"
                  className="read-more-link text-decoration-none dg-blog-btn-text"
                >
                  Explore More <i className="fas fa-arrow-right ms-2"></i>
                </a>
              </article>
            </div>
          </div>
        </div>
      </section>
    </>
  );
};

export default BlogItems;
