import { useState } from "react";
import axios from "axios";

import { Container, Form, Button, FloatingLabel, Alert } from "react-bootstrap";

const CreateEvent = ({ onEventCreated, id }) => {
    const [errMsg, setErrMsg] = useState("");
    const [createEventForm, setCreateEventForm] = useState({
        title: "",
        description: "",
        startTime: "",
        endTime: "",
        group_id: +id,
    });

    const handleChange = (event) => {
        setErrMsg("");
        const { name, value } = event?.target;
        setCreateEventForm((prevState) => ({
            ...prevState,
            [name]: value,
        }));
    };

    const handleSubmit = async (event) => {
        event.preventDefault();

        if (
            createEventForm?.title === "" ||
            createEventForm?.description === "" ||
            createEventForm?.startTime === ""
        ) {
            setErrMsg("Some fields missing");
            return;
        }
        // Send the createEventForm data to the backend handler
        try {
            const response = await axios.post(
                "http://localhost:8000/creategroupevent",
                JSON.stringify({
                    ...createEventForm,
                    startTime: new Date(
                        createEventForm.startTime
                    ).toISOString(),
                    endTime: new Date(createEventForm.endTime).toISOString(),
                }),
                { withCredentials: true },
                {
                    headers: { "Content-Type": "application/json" },
                }
            );

            if (response.status === 200) {
                setErrMsg("New Event Created");
                onEventCreated();
            } else {
                setErrMsg("Server responded with an unexpected status code");
            }
        } catch (err) {}
    };

    return (
        <>
            {errMsg && (
                <Alert
                    variant={
                        errMsg === "New Event Created" ? "success" : "danger"
                    }
                    className="text-center"
                >
                    {errMsg}
                </Alert>
            )}
            <Form onSubmit={handleSubmit}>
                {errMsg && <p className="error">{errMsg}</p>}
                <form className="pop-form">
                    Title:
                    <label className="input-big">
                        <input
                            type="text"
                            name="title"
                            value={createEventForm.title}
                            onChange={handleChange}
                            required
                        />
                    </label>
                    <br />
                    Description:
                    <label>
                        <textarea
                            className="text-big"
                            name="description"
                            value={createEventForm.description}
                            onChange={handleChange}
                            required
                        ></textarea>
                    </label>
                    <br />
                    Start Time:
                    <label>
                        <input
                            type="datetime-local"
                            name="startTime"
                            value={createEventForm.startTime}
                            onChange={handleChange}
                            required
                        />
                    </label>
                    <br />
                    End Time:
                    <label>
                        <input
                            type="datetime-local"
                            name="endTime"
                            value={createEventForm.endTime}
                            onChange={handleChange}
                            required
                        />
                    </label>
                    <br />
                </form>
                <Button type="submit">Create</Button>
            </Form>
        </>
    );
};

export default CreateEvent;
