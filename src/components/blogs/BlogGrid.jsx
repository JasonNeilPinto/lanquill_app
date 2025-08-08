import React from "react";
import { Link } from "react-router-dom";
// import BlogPagination from "../blogs/BlogPagination";

const BlogGrid = () => {
  return (
    <>
      <section className="masonary-blog-section ptb-120">
        <div className="container">
          <div className="row">
            <div className="col-lg-6 col-md-12">
              <div className="single-article feature-article rounded-custom my-3">
                <Link to="/blog-single" className="article-img">
                  <img
                    src="/img/blog/beginnerit.png"
                    alt="article"
                    className="img-fluid"
                  />
                </Link>
                <div className="article-content p-4">
                  <div className="article-category mb-4 d-block">
                    <Link
                      to="#"
                      className="d-inline-block text-dark badge bg-primary-soft"
                    >
                      Blog
                    </Link>
                  </div>
                  <Link to="/blog-single">
                    <h2 className="h5 article-title limit-2-line-text">
                      A Beginner’s Guide to IT Outsourcing in Bangalore
                    </h2>
                  </Link>
                  <p className="limit-2-line-text">
                    Bangalore, often referred to as the Silicon Valley of India,
                    is a global hub for technology and innovation. With its
                    thriving IT ecosystem, the city has become a hotspot for
                    businesses looking to outsource their IT needs. Whether
                    you’re a startup, SME, or large enterprise, IT outsourcing
                    can help you reduce costs, access specialized expertise, and
                    focus on your core business operations. If you’re new to IT
                    outsourcing, this guide will walk you through everything you
                    need to know to get started.
                  </p>

                  <Link to="#">
                    <div className="d-flex align-items-center pt-4">
                      <div className="avatar">
                        <img
                          src="/img/testimonial/1.jpg"
                          alt="avatar"
                          width="40"
                          className="img-fluid rounded-circle me-3"
                        />
                      </div>
                      <div className="avatar-info">
                        <h6 className="mb-0 avatar-name">Donna Martin</h6>
                        <span className="small fw-medium text-muted">
                          April 24, 2022
                        </span>
                      </div>
                    </div>
                  </Link>
                </div>
              </div>
            </div>
            <div className="col-lg-6 col-md-12">
              <div className="single-article feature-article rounded-custom my-3">
                <Link to="/blog-single" className="article-img">
                  <img
                    src="/img/blog/Talent.png"
                    alt="article"
                    className="img-fluid"
                  />
                </Link>
                <div className="article-content p-4">
                  <div className="article-category mb-4 d-block">
                    <Link
                      to="#"
                      className="d-inline-block text-dark badge bg-primary-soft"
                    >
                      Blog
                    </Link>
                  </div>
                  <Link to="/blog-second">
                    <h2 className="h5 article-title limit-2-line-text">
                      Talent Acquisition Trends 2025 fit
                    </h2>
                  </Link>
                  <p className="limit-2-line-text">
                    As we approach 2025, talent acquisition is poised to undergo
                    transformative changes driven by technology, evolving
                    workforce expectations, and global economic shifts. Talent
                    managers are discovering that AI need human control to have
                    a significant impact on hiring issues. Anticipate a more
                    balanced and judicious approach to AI in talent recruiting
                    by 2025.
                  </p>

                  <Link to="#">
                    <div className="d-flex align-items-center pt-4">
                      <div className="avatar">
                        <img
                          src="/img/testimonial/4.jpg"
                          alt="avatar"
                          width="40"
                          className="img-fluid rounded-circle me-3"
                        />
                      </div>
                      <div className="avatar-info">
                        <h6 className="mb-0 avatar-name">Donna Martin</h6>
                        <span className="small fw-medium text-muted">
                          April 24, 2022
                        </span>
                      </div>
                    </div>
                  </Link>
                </div>
              </div>
            </div>
          </div>
          <div className="row">
            <div className="col-lg-4 col-md-6">
              <div className="single-article rounded-custom my-3">
                <Link to="/blog-single" className="article-img">
                  <img
                    src="/img/blog/big-ideas.png"
                    alt="article"
                    className="img-fluid"
                  />
                </Link>
                <div className="article-content p-4">
                  <div className="article-category mb-4 d-block">
                    <Link
                      to="#"
                      className="d-inline-block text-dark badge bg-primary-soft"
                    >
                      Blogs
                    </Link>
                  </div>
                  <Link to="/blog-ShapingIndia">
                    <h2 className="h5 article-title limit-2-line-text">
                      Shaping India's Future: 12 Big Ideas for 2025
                    </h2>
                  </Link>
                  <p className="limit-2-line-text">
                    As India continues its upward trajectory, it's imperative
                    for business leaders to stay ahead of the curve and
                    anticipate the trends that will shape the country's future.
                    We've identified 15 big ideas that will have a significant
                    impact on India's growth and development in 2025.
                  </p>

                  <Link to="#">
                    <div className="d-flex align-items-center pt-4">
                      <div className="avatar">
                        <img
                          src="/img/testimonial/6.jpg"
                          alt="avatar"
                          width="40"
                          className="img-fluid rounded-circle me-3"
                        />
                      </div>
                      <div className="avatar-info">
                        <h6 className="mb-0 avatar-name">Jane Martin</h6>
                        <span className="small fw-medium text-muted">
                          April 24, 2021
                        </span>
                      </div>
                    </div>
                  </Link>
                </div>
              </div>
            </div>
            <div className="col-lg-4 col-md-6">
              <div className="single-article rounded-custom my-3">
                <Link to="/blog-single" className="article-img">
                  <img
                    src="/img/blog/war.png"
                    alt="article"
                    className="img-fluid"
                  />
                </Link>
                <div className="article-content p-4">
                  <div className="article-category mb-4 d-block">
                    <Link
                      to="#"
                      className="d-inline-block text-dark badge bg-primary-soft"
                    >
                      Blog
                    </Link>
                  </div>
                  <Link to="/blog-theWarfortalent">
                    <h2 className="h5 article-title limit-2-line-text">
                      The War for Talent: How IT Staffing Can Be Your Secret
                      Weapon
                    </h2>
                  </Link>
                  <p className="limit-2-line-text">
                    The tech industry is on fire. From artificial intelligence
                    to cybersecurity, innovation is exploding, and businesses
                    are scrambling to keep pace. But with this rapid growth
                    comes fierce competition for skilled IT professionals. It's
                    a full-blown war for talent, and companies without a
                    strategic plan are at risk of falling behind.
                  </p>

                  <Link to="#">
                    <div className="d-flex align-items-center pt-4">
                      <div className="avatar">
                        <img
                          src="/img/testimonial/1.jpg"
                          alt="avatar"
                          width="40"
                          className="img-fluid rounded-circle me-3"
                        />
                      </div>
                      <div className="avatar-info">
                        <h6 className="mb-0 avatar-name">Veronica P. Byrd</h6>
                        <span className="small fw-medium text-muted">
                          April 24, 2021
                        </span>
                      </div>
                    </div>
                  </Link>
                </div>
              </div>
            </div>
            <div className="col-lg-4 col-md-6">
              <div className="single-article rounded-custom my-3">
                <Link to="/blog-single" className="article-img">
                  <img
                    src="/img/blog/Generative_AI.png"
                    alt="article"
                    className="img-fluid"
                  />
                </Link>
                <div className="article-content p-4">
                  <div className="article-category mb-4 d-block">
                    <Link
                      to="#"
                      className="d-inline-block text-dark badge bg-primary-soft"
                    >
                      Blog
                    </Link>
                  </div>
                  <Link to="/blog-generativeai">
                    <h2 className="h5 article-title limit-2-line-text">
                      How Generative AI Will Transform the Legal Industry
                    </h2>
                  </Link>
                  <p className="limit-2-line-text">
                    The advent of artificial intelligence (AI) is heralding a
                    new era of transformation across various industries, and the
                    legal sector is no exception. Generative AI, a subset of AI
                    that focuses on creating new content and ideas, is poised to
                    revolutionize the way legal services are delivered,
                    fundamentally altering the traditional landscape of the
                    legal profession. This blog explores how generative AI will
                    affect the legal industry, from enhancing efficiency and
                    reducing costs to posing ethical challenges and reshaping
                    the roles of legal professionals.
                  </p>

                  <Link to="#">
                    <div className="d-flex align-items-center pt-4">
                      <div className="avatar">
                        <img
                          src="/img/testimonial/3.jpg"
                          alt="avatar"
                          width="40"
                          className="img-fluid rounded-circle me-3"
                        />
                      </div>
                      <div className="avatar-info">
                        <h6 className="mb-0 avatar-name">Martin Gilbert</h6>
                        <span className="small fw-medium text-muted">
                          April 24, 2021
                        </span>
                      </div>
                    </div>
                  </Link>
                </div>
              </div>
            </div>
            <div className="col-lg-4 col-md-6">
              <div className="single-article rounded-custom my-3">
                <Link to="/blog-single" className="article-img">
                  <img
                    src="/img/blog/genAINew.jpg"
                    alt="article"
                    className="img-fluid"
                  />
                </Link>
                <div className="article-content p-4">
                  <div className="article-category mb-4 d-block">
                    <Link
                      to="#"
                      className="d-inline-block text-dark badge bg-primary-soft"
                    >
                      Blog
                    </Link>
                  </div>
                  <Link to="/blog-emergenceofgenai">
                    <h2 className="h5 article-title limit-2-line-text">
                      Emergence of GenAI ChatBots in the Healthcare Sector: A
                      Comprehensive Analysis{" "}
                    </h2>
                  </Link>
                  <p className="limit-2-line-text">
                    The healthcare sector is experiencing a digital revolution,
                    with the integration of artificial intelligence (AI) playing
                    a pivotal role. Among the various AI applications,
                    generative AI (GenAI) chatbots have emerged as a significant
                    advancement. This article delves into the emergence of GenAI
                    chatbots in healthcare, comparing them with traditional
                    chatbots, and examining the underlying technology that sets
                    them apart. We will also highlight recent research reports
                    from reputable institutions such as McKinsey, BCG, and
                    Harvard University to provide a comprehensive analysis.
                    Additionally, we will include a comparison table and
                    real-world examples of how GenAI chatbots respond to
                    customer queries compared to traditional chatbots.
                  </p>

                  <Link to="#">
                    <div className="d-flex align-items-center pt-4">
                      <div className="avatar">
                        <img
                          src="/img/testimonial/4.jpg"
                          alt="avatar"
                          width="40"
                          className="img-fluid rounded-circle me-3"
                        />
                      </div>
                      <div className="avatar-info">
                        <h6 className="mb-0 avatar-name">Raymond H. Martin</h6>
                        <span className="small fw-medium text-muted">
                          May 4, 2021
                        </span>
                      </div>
                    </div>
                  </Link>
                </div>
              </div>
            </div>
            <div className="col-lg-4 col-md-6">
              <div className="single-article rounded-custom my-3">
                <Link to="/blog-single" className="article-img">
                  <img
                    src="/img/blog/howgenai.png"
                    alt="article"
                    className="img-fluid"
                  />
                </Link>
                <div className="article-content p-4">
                  <div className="article-category mb-4 d-block">
                    <Link
                      to="#"
                      className="d-inline-block text-dark badge bg-primary-soft"
                    >
                      Blog
                    </Link>
                  </div>
                  <Link to="/blog-genairevolutionizes">
                    <h2 className="h5 article-title limit-2-line-text">
                      How Gen AI Revolutionizes Coding: The Future of AI-Powered
                      Coding Assistants
                    </h2>
                  </Link>
                  <p className="limit-2-line-text">
                    In the fast-paced world of software development, efficiency
                    and innovation are paramount. Generative AI (Gen AI) is
                    rapidly emerging as a game-changer in the realm of coding.
                    By automating code generation and providing intelligent
                    coding assistance, Gen AI is not only enhancing productivity
                    but also significantly reducing costs for companies. This
                    blog explores the transformative impact of Gen AI on coding,
                    supported by research studies and practical examples.
                  </p>

                  <Link to="#">
                    <div className="d-flex align-items-center pt-4">
                      <div className="avatar">
                        <img
                          src="/img/testimonial/5.jpg"
                          alt="avatar"
                          width="40"
                          className="img-fluid rounded-circle me-3"
                        />
                      </div>
                      <div className="avatar-info">
                        <h6 className="mb-0 avatar-name">Luthar Martin</h6>
                        <span className="small fw-medium text-muted">
                          Jan 24, 2021
                        </span>
                      </div>
                    </div>
                  </Link>
                </div>
              </div>
            </div>
            <div className="col-lg-4 col-md-6">
              <div className="single-article rounded-custom my-3">
                <Link to="/blog-single" className="article-img">
                  <img
                    src="/img/blog/ecommerce.png"
                    alt="article"
                    className="img-fluid"
                  />
                </Link>
                <div className="article-content p-4">
                  <div className="article-category mb-4 d-block">
                    <Link
                      to="#"
                      className="d-inline-block text-dark badge bg-primary-soft"
                    >
                      Blog
                    </Link>
                  </div>
                  <Link to="/blog-ecommerce">
                    <h2 className="h5 article-title limit-2-line-text">
                      Cybersecurity Best Practices for Protecting Customer Data
                      in E-commerce: A Guide for CTOs
                    </h2>
                  </Link>
                  <p className="limit-2-line-text">
                    In today's digital age, where e-commerce transactions have
                    become the norm, protecting customer data is paramount for
                    maintaining trust and credibility. With cyber threats on the
                    rise, Chief Technology Officers (CTOs) in the e-commerce
                    industry need to implement robust cybersecurity measures to
                    safeguard sensitive customer information. In this
                    comprehensive guide, we'll delve into best practices and
                    provide a step-by-step implementation plan to fortify your
                    e-commerce platform against potential data breaches.
                  </p>

                  <Link to="#">
                    <div className="d-flex align-items-center pt-4">
                      <div className="avatar">
                        <img
                          src="/img/testimonial/6.jpg"
                          alt="avatar"
                          width="40"
                          className="img-fluid rounded-circle me-3"
                        />
                      </div>
                      <div className="avatar-info">
                        <h6 className="mb-0 avatar-name">Donna Martin</h6>
                        <span className="small fw-medium text-muted">
                          April 24, 2022
                        </span>
                      </div>
                    </div>
                  </Link>
                </div>
              </div>
            </div>
            {/* <div className="col-lg-4 col-md-6">
              <div className="single-article rounded-custom my-3">
                <Link to="/blog-single" className="article-img">
                  <img
                    src="/img/blog/blog-7.jpg"
                    alt="article"
                    className="img-fluid"
                  />
                </Link>
                <div className="article-content p-4">
                  <div className="article-category mb-4 d-block">
                    <Link
                      to="#"
                      className="d-inline-block text-dark badge bg-danger-soft"
                    >
                      Design
                    </Link>
                  </div>
                  <Link to="/blog-single">
                    <h2 className="h5 article-title limit-2-line-text">
                      New analyst report: Digital product management tools and
                      tech
                    </h2>
                  </Link>
                  <p className="limit-2-line-text">
                    Society is fragmenting into two parallel realities. In one
                    reality, you have infinite upside and opportunity. In the
                    other reality, you’ll continue to see the gap between your
                    standard of living and those at the top grow more and more.
                  </p>

                  <Link to="#">
                    <div className="d-flex align-items-center pt-4">
                      <div className="avatar">
                        <img
                          src="/img/testimonial/1.jpg"
                          alt="avatar"
                          width="40"
                          className="img-fluid rounded-circle me-3"
                        />
                      </div>
                      <div className="avatar-info">
                        <h6 className="mb-0 avatar-name">Donna R. Book</h6>
                        <span className="small fw-medium text-muted">
                          April 24, 2021
                        </span>
                      </div>
                    </div>
                  </Link>
                </div>
              </div>
            </div> */}
            {/* <div className="col-lg-4 col-md-6">
              <div className="single-article rounded-custom my-3">
                <Link to="/blog-single" className="article-img">
                  <img
                    src="/img/blog/blog-8.jpg"
                    alt="article"
                    className="img-fluid"
                  />
                </Link>
                <div className="article-content p-4">
                  <div className="article-category mb-4 d-block">
                    <Link
                      to="#"
                      className="d-inline-block text-dark badge bg-primary-soft"
                    >
                      Development
                    </Link>
                  </div>
                  <Link to="/blog-single">
                    <h2 className="h5 article-title limit-2-line-text">
                      A frank discussion about diversity, inclusion, and
                      allyship
                    </h2>
                  </Link>
                  <p className="limit-2-line-text">
                    Society is fragmenting into two parallel realities. In one
                    reality, you have infinite upside and opportunity. In the
                    other reality, you’ll continue to see the gap between your
                    standard of living and those at the top grow more and more.
                  </p>

                  <Link to="#">
                    <div className="d-flex align-items-center pt-4">
                      <div className="avatar">
                        <img
                          src="/img/testimonial/3.jpg"
                          alt="avatar"
                          width="40"
                          className="img-fluid rounded-circle me-3"
                        />
                      </div>
                      <div className="avatar-info">
                        <h6 className="mb-0 avatar-name">Donna R. Martin</h6>
                        <span className="small fw-medium text-muted">
                          April 24, 2021
                        </span>
                      </div>
                    </div>
                  </Link>
                </div>
              </div>
            </div> */}
            {/* <div className="col-lg-4 col-md-6">
              <div className="single-article rounded-custom my-3">
                <Link to="/blog-single" className="article-img">
                  <img
                    src="/img/blog/blog-9.jpg"
                    alt="article"
                    className="img-fluid"
                  />
                </Link>
                <div className="article-content p-4">
                  <div className="article-category mb-4 d-block">
                    <Link
                      to="#"
                      className="d-inline-block text-dark badge bg-warning-soft"
                    >
                      Design
                    </Link>
                  </div>
                  <Link to="/blog-single">
                    <h2 className="h5 article-title limit-2-line-text">
                      4 steps for measuring the impact of product discovery
                    </h2>
                  </Link>
                  <p className="limit-2-line-text">
                    Society is fragmenting into two parallel realities. In one
                    reality, you have infinite upside and opportunity. In the
                    other reality, you’ll continue to see the gap between your
                    standard of living and those at the top grow more and more.
                  </p>

                  <Link to="#">
                    <div className="d-flex align-items-center pt-4">
                      <div className="avatar">
                        <img
                          src="/img/testimonial/2.jpg"
                          alt="avatar"
                          width="40"
                          className="img-fluid rounded-circle me-3"
                        />
                      </div>
                      <div className="avatar-info">
                        <h6 className="mb-0 avatar-name">Martin Luthar</h6>
                        <span className="small fw-medium text-muted">
                          April 24, 2021
                        </span>
                      </div>
                    </div>
                  </Link>
                </div>
              </div>
            </div> */}
          </div>
          {/* <BlogPagination /> */}
        </div>
      </section>
    </>
  );
};

export default BlogGrid;
