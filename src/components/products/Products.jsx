import React from "react";
const Products = () => {
  return (
    <>
      <div
        id="div-products"
        className="social-media-section section-space--sm bgc-background"
      >
        <div className="section-space--sm-bottom">
          <div className="container">
            <div className="row justify-content-center">
              <div className="col-md-10 col-xl-8">
                <h3 className="heading-3 clr-text font-weight-semibold text-center mb-0">
                  Our Cutting-Edge {" "}
                  <span className="clr-blue">Products and Platforms</span>
                </h3>
              </div>
            </div>
          </div>
        </div>
        <div className="container">
          <div className="row">
            <div className="col-12">
              <div className="list-group flex-row justify-content-center overflow-auto margin-bottom-13">
                <a
                  className="social-media-btn active"
                  data-bs-toggle="list"
                  href="#pramitihr"
                >
                  <span className="d-grid place-content-center width-12 height-12 rounded-circle bgc-text-1 clr-paragraph flex-shrink-0 fs-16">
                    <i className="pramiti-icon"></i>
                  </span>
                  <span className="d-inline-block font-weight-bold fs-16 clr-paragraph">
                    PramitiHR
                  </span>
                </a>
                <a
                  className="social-media-btn"
                  data-bs-toggle="list"
                  href="#lanquill"
                >
                  <span className="d-grid place-content-center width-12 height-12 rounded-circle bgc-text-1 clr-paragraph flex-shrink-0 fs-16">
                    <i className="lanquill-icon"></i>
                  </span>
                  <span className="d-inline-block font-weight-bold fs-16 clr-paragraph">
                    Lanquill
                  </span>
                </a>
                <a
                  className="social-media-btn"
                  data-bs-toggle="list"
                  href="#genAI"
                >
                  <span className="d-grid place-content-center width-12 height-12 rounded-circle bgc-text-1 clr-paragraph flex-shrink-0 fs-16">
                    <i className="openAi-icon"></i>
                  </span>
                  <span className="d-inline-block font-weight-bold fs-16 clr-paragraph">
                    Gen AI Sandbox
                  </span>
                </a>
                <a
                  className="social-media-btn"
                  data-bs-toggle="list"
                  href="#ecommerce"
                >
                  <span className="d-grid place-content-center width-12 height-12 rounded-circle bgc-text-1 clr-paragraph flex-shrink-0 fs-16">
                    <i
                      className="far fa-cart-shopping"
                      style={{ color: "#46a4ec" }}
                    ></i>
                  </span>
                  <span className="d-inline-block font-weight-bold fs-16 clr-paragraph">
                    Reverse Auction e-commerce
                  </span>
                </a>
                <a
                  className="social-media-btn"
                  data-bs-toggle="list"
                  href="#entertainment"
                >
                  <span className="d-grid place-content-center width-12 height-12 rounded-circle bgc-text-1 clr-paragraph flex-shrink-0 fs-16">
                    <i
                      className="fa-solid fa-clapperboard"
                      style={{ color: "#7a83c2" }}
                    ></i>
                  </span>
                  <span className="d-inline-block font-weight-bold fs-16 clr-paragraph">
                    OnDemand Entertainment
                  </span>
                </a>
              </div>
              <div className="tab-content">
                <div className="tab-pane fade show active" id="pramitihr">
                  <div className="border border-blue-clr rounded-4 padding-4">
                    <div className="bgc-white rounded-4 section-space padding-x-2 padding-x-xsm-4 padding-x-sm-8 padding-x-md-12 padding-end-xxl-25">
                      <div className="row">
                        <div className="col-lg-5 col-xl-6 col-xxl-7">
                          <img
                            src="/img/products/pramiti-product.webp"
                            alt="PramitiHR"
                            className="img-fluid rounded-4"
                          />
                        </div>
                        <div className="col-lg-7 col-xl-6 col-xxl-5 ms-lg-auto">
                          <h5 className="heading-5 font-weight-semibold margin-bottom-5 clr-text">
                            PramitiHR: Hire Smarter with Generative AI
                            Interviews
                          </h5>
                          <p className="clr-paragraph margin-bottom-8">
                            Automate screenings and interviews, reducing
                            time-to-hire and enabling hiring managers to focus
                            on top candidates.
                          </p>
                          <ul className="list gap-3">
                            <li>
                              <div className="d-flex gap-2 align-items-start clr-paragraph">
                                <div className="flex-shrink-0">
                                  <i className="fas fa-check-circle"></i>
                                </div>
                                <p className="mb-0 flex-grow-1">
                                  <b>Efficiency:</b> Faster shortlisting and
                                  decision making
                                </p>
                              </div>
                            </li>
                            <li>
                              <div className="d-flex gap-2 align-items-start clr-paragraph">
                                <div className="flex-shrink-0">
                                  <i className="fas fa-check-circle"></i>
                                </div>
                                <p className="mb-0 flex-grow-1">
                                  <b>Bias Reduction:</b> Promote diversity and
                                  fairer assessments
                                </p>
                              </div>
                            </li>
                            <li>
                              <div className="d-flex gap-2 align-items-start clr-paragraph">
                                <div className="flex-shrink-0">
                                  <i className="fas fa-check-circle"></i>
                                </div>
                                <p className="mb-0 flex-grow-1">
                                  <b>Scalability:</b> Handle large volumes of
                                  interviews simultaneously
                                </p>
                              </div>
                            </li>
                          </ul>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div className="tab-pane fade" id="lanquill">
                  <div className="border border-blue-clr rounded-4 padding-4">
                    <div className="lanquill-content bgc-white rounded-4 section-space padding-x-2 padding-x-xsm-4 padding-x-sm-8 padding-x-md-12 padding-end-xxl-25">
                      <div className="row">
                        <div className="col-lg-5 col-xl-6 col-xxl-7">
                          <img
                            src="/img/products/lanquill-product.webp"
                            alt="Lanquill"
                            className="img-fluid rounded-4"
                          />
                        </div>
                        <div className="col-lg-7 col-xl-6 col-xxl-5 ms-lg-auto">
                          <h5 className="heading-5 font-weight-semibold margin-bottom-5 clr-text">
                            Lanquill: An AI-powered English language learning
                            lab
                          </h5>
                          <p className="clr-paragraph margin-bottom-8">
                            Designed to develop communication skills through
                            listening, speaking, reading, writing, vocabulary
                            and grammar practices. Get CEFR align certificates
                            and access to huge collection of personalized
                            learning content.
                          </p>
                          <ul className="list gap-3">
                            <li>
                              <div className="d-flex gap-2 align-items-center clr-paragraph">
                                <div className="flex-shrink-0">
                                  <i className="fas fa-check-circle"></i>
                                </div>
                                <p className="mb-0 flex-grow-1">
                                  Practice lab (LSRWGV)
                                </p>
                              </div>
                            </li>
                            <li>
                              <div className="d-flex gap-2 align-items-center clr-paragraph">
                                <div className="flex-shrink-0">
                                  <i className="fas fa-check-circle"></i>
                                </div>
                                <p className="mb-0 flex-grow-1">
                                  CEFR Aligned Certifications (Pre A1 - C2)
                                </p>
                              </div>
                            </li>
                            <li>
                              <div className="d-flex gap-2 align-items-center clr-paragraph">
                                <div className="flex-shrink-0">
                                  <i className="fas fa-check-circle"></i>
                                </div>
                                <p className="mb-0 flex-grow-1">
                                  Personalized Library (700+ Hours content)
                                </p>
                              </div>
                            </li>
                          </ul>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div className="tab-pane fade" id="genAI">
                  <div className="border border-blue-clr rounded-4 padding-4">
                    <div className="bgc-white rounded-4 section-space padding-x-2 padding-x-xsm-4 padding-x-sm-8 padding-x-md-12 padding-end-xxl-25">
                      <div className="row">
                        <div className="col-lg-5 col-xl-6 col-xxl-7">
                          <img
                            src="/img/products/sandbox-product.webp"
                            alt="Lanquill"
                            className="img-fluid rounded-4"
                          />
                        </div>
                        <div className="col-lg-7 col-xl-6 col-xxl-5 ms-lg-auto">
                          <h5 className="heading-5 font-weight-semibold margin-bottom-5 clr-text">
                            Gen AI Sandbox
                          </h5>
                          <p className="clr-paragraph margin-bottom-8">
                            A secure, flexible environment to prototype and
                            deploy Generative AI solutions tailored to your
                            business needs.
                          </p>
                          <ul className="list gap-3">
                            <li>
                              <div className="d-flex gap-2 align-items-center clr-paragraph">
                                <div className="flex-shrink-0">
                                  <i className="fas fa-check-circle"></i>
                                </div>
                                <p className="mb-0 flex-grow-1">
                                  Rapid prototyping for AI-driven solutions
                                </p>
                              </div>
                            </li>
                            <li>
                              <div className="d-flex gap-2 align-items-center clr-paragraph">
                                <div className="flex-shrink-0">
                                  <i className="fas fa-check-circle"></i>
                                </div>
                                <p className="mb-0 flex-grow-1">
                                  Secure and scalable testing environment
                                </p>
                              </div>
                            </li>
                            <li>
                              <div className="d-flex gap-2 align-items-center clr-paragraph">
                                <div className="flex-shrink-0">
                                  <i className="fas fa-check-circle"></i>
                                </div>
                                <p className="mb-0 flex-grow-1">
                                  Tailored to your unique business needs
                                </p>
                              </div>
                            </li>
                          </ul>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div className="tab-pane fade" id="ecommerce">
                  <div className="border border-blue-clr rounded-4 padding-4">
                    <div className="bgc-white rounded-4 section-space padding-x-2 padding-x-xsm-4 padding-x-sm-8 padding-x-md-12 padding-end-xxl-25">
                      <div className="row">
                        <div className="col-lg-5 col-xl-6 col-xxl-7">
                          <img
                            src="/img/products/e-commerce-product.webp"
                            alt="Lanquill"
                            className="img-fluid rounded-4"
                          />
                        </div>
                        <div className="col-lg-7 col-xl-6 col-xxl-5 ms-lg-auto">
                          <h5 className="heading-5 font-weight-semibold margin-bottom-5 clr-text">
                            Reverse Auction e-Commerce
                          </h5>
                          <p className="clr-paragraph margin-bottom-8">
                            A reverse auction e-commerce platform for consumers
                            to buy at the best price, ensuring the best deals as
                            sellers bid to offer you the lowest price.
                          </p>
                          <ul className="list gap-3">
                            <li>
                              <div className="d-flex gap-2 align-items-center clr-paragraph">
                                <div className="flex-shrink-0">
                                  <i className="fas fa-check-circle"></i>
                                </div>
                                <p className="mb-0 flex-grow-1">
                                  Competitive bidding for the best prices
                                </p>
                              </div>
                            </li>
                            <li>
                              <div className="d-flex gap-2 align-items-center clr-paragraph">
                                <div className="flex-shrink-0">
                                  <i className="fas fa-check-circle"></i>
                                </div>
                                <p className="mb-0 flex-grow-1">
                                  Increased savings with transparent deals
                                </p>
                              </div>
                            </li>
                            <li>
                              <div className="d-flex gap-2 align-items-center clr-paragraph">
                                <div className="flex-shrink-0">
                                  <i className="fas fa-check-circle"></i>
                                </div>
                                <p className="mb-0 flex-grow-1">
                                  User-friendly platform for seamless
                                  transactions
                                </p>
                              </div>
                            </li>
                          </ul>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div className="tab-pane fade" id="entertainment">
                  <div className="border border-blue-clr rounded-4 padding-4">
                    <div className="bgc-white rounded-4 section-space padding-x-2 padding-x-xsm-4 padding-x-sm-8 padding-x-md-12 padding-end-xxl-25">
                      <div className="row">
                        <div className="col-lg-5 col-xl-6 col-xxl-7">
                          <img
                            src="/img/products/entertainment-product.webp"
                            alt="Lanquill"
                            className="img-fluid rounded-4"
                          />
                        </div>
                        <div className="col-lg-7 col-xl-6 col-xxl-5 ms-lg-auto">
                          <h5 className="heading-5 font-weight-semibold margin-bottom-5 clr-text">
                            OnDemand Entertainment
                          </h5>
                          <p className="clr-paragraph margin-bottom-8">
                            Deliver unmatched passenger experiences with our
                            OnDemand Entertainment solutions, offering movies,
                            music, games, and more at their fingertips. Built
                            for speed, reliability, and engagement, it keeps
                            users entertained anytime, anywhere.
                          </p>
                          <ul className="list gap-3">
                            <li>
                              <div className="d-flex gap-2 align-items-center clr-paragraph">
                                <div className="flex-shrink-0">
                                  <i className="fas fa-check-circle"></i>
                                </div>
                                <p className="mb-0 flex-grow-1">
                                  Wide range of on-demand media content
                                </p>
                              </div>
                            </li>
                            <li>
                              <div className="d-flex gap-2 align-items-center clr-paragraph">
                                <div className="flex-shrink-0">
                                  <i className="fas fa-check-circle"></i>
                                </div>
                                <p className="mb-0 flex-grow-1">
                                  Seamless integration across devices
                                </p>
                              </div>
                            </li>
                            <li>
                              <div className="d-flex gap-2 align-items-center clr-paragraph">
                                <div className="flex-shrink-0">
                                  <i className="fas fa-check-circle"></i>
                                </div>
                                <p className="mb-0 flex-grow-1">
                                  Optimized for high-quality streaming
                                </p>
                              </div>
                            </li>
                          </ul>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default Products;
