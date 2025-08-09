import React, { useState } from "react";

const ContactFormTwo = () => {
  const [formData, setFormData] = useState({
    firstName: "",
    lastName: "",
    mobile: "",
    email: "",
    message: "",
  });

  const [isSubmitting, setIsSubmitting] = useState(false);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");

  const handleChange = (e) => {
    const { id, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [id]: value,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setIsSubmitting(true);
    setError("");
    setSuccess("");

    const payload = {
      name: `${formData.firstName} ${formData.lastName}`.trim(),
      mobile: formData.mobile,
      email: formData.email,
      message: formData.message,
    };

    try {
      const response = await fetch("https://lanquill.com/auth/contact-us-na", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(payload),
      });

      if (!response.ok) {
        throw new Error("Something went wrong!");
      }

      setSuccess("Message sent successfully!");
      setFormData({
        firstName: "",
        lastName: "",
        mobile: "",
        email: "",
        message: "",
      });
    } catch (err) {
      console.log("Error", err);
      setError("Failed to send message. Please try again later.");
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <>
      <section
        className="contact-us-form pt-60 pb-120"
        style={{
          background:
            "url('/img/shape/contact-us-bg.svg')no-repeat center bottom",
        }}
      >
        <div className="container">
          <div className="row justify-content-lg-between align-items-center">
            <div className="col-lg-6 col-md-8">
              <div className="section-heading">
                <h2>Talk to Our Sales & Marketing Department Team</h2>
                <p>
                  Collaboratively promote client-focused convergence vis-a-vis
                  customer directed alignments via standardized infrastructures.
                </p>
              </div>
              <form onSubmit={handleSubmit} className="register-form">
                <div className="row">
                  <div className="col-sm-6">
                    <label htmlFor="firstName" className="mb-1">
                      First name <span className="text-danger">*</span>
                    </label>
                    <div className="input-group mb-3">
                      <input
                        type="text"
                        className="form-control"
                        id="firstName"
                        required
                        placeholder="First name"
                        value={formData.firstName}
                        onChange={handleChange}
                      />
                    </div>
                  </div>
                  <div className="col-sm-6">
                    <label htmlFor="lastName" className="mb-1">
                      Last name
                    </label>
                    <div className="input-group mb-3">
                      <input
                        type="text"
                        className="form-control"
                        id="lastName"
                        placeholder="Last name"
                        value={formData.lastName}
                        onChange={handleChange}
                      />
                    </div>
                  </div>
                  <div className="col-sm-6">
                    <label htmlFor="mobile" className="mb-1">
                      Phone <span className="text-danger">*</span>
                    </label>
                    <div className="input-group mb-3">
                      <input
                        type="text"
                        className="form-control"
                        id="mobile"
                        required
                        placeholder="Phone"
                        value={formData.mobile}
                        onChange={handleChange}
                      />
                    </div>
                  </div>
                  <div className="col-sm-6">
                    <label htmlFor="email" className="mb-1">
                      Email <span className="text-danger">*</span>
                    </label>
                    <div className="input-group mb-3">
                      <input
                        type="email"
                        className="form-control"
                        id="email"
                        required
                        placeholder="Email"
                        value={formData.email}
                        onChange={handleChange}
                      />
                    </div>
                  </div>
                  <div className="col-12">
                    <label htmlFor="message" className="mb-1">
                      Message <span className="text-danger">*</span>
                    </label>
                    <div className="input-group mb-3">
                      <textarea
                        className="form-control"
                        id="message"
                        required
                        placeholder="How can we help you?"
                        value={formData.message}
                        onChange={handleChange}
                        style={{ height: "120px" }}
                      ></textarea>
                    </div>
                  </div>
                </div>
                <button
                  type="submit"
                  className="btn btn-primary mt-4"
                  disabled={isSubmitting}
                >
                  {isSubmitting ? "Sending..." : "Get in Touch"}
                </button>

                {error && <p className="text-danger mt-2">{error}</p>}
                {success && <p className="text-success mt-2">{success}</p>}
              </form>
            </div>
            <div className="col-lg-5 col-md-10">
              <div className="contact-us-img">
                <img
                  src="/img/contact-us-img-2.svg"
                  alt="contact us"
                  className="img-fluid"
                />
              </div>
            </div>
          </div>
        </div>
      </section>
    </>
  );
};

export default ContactFormTwo;
