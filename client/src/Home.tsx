import { useState } from "react";
import { Collapse } from "./components/Collapse";
import { Modal } from "./components/Modal";
import { QuickSlotGenerator, useQuickSlotGenerator } from "./components/QuickSlotGenerator";
import { TimeSlotForm, useTimeSlotModal } from "./components/TimeSlotForm";
import { TimeSlotList } from "./components/TimeSlotList";
import type { TimeSlot } from "./types";

export const Home = () => {
  const [timeSlots, setTimeSlots] = useState<TimeSlot[]>([]);

  const {
    isModalOpen,
    modalData,
    setModalData,
    editingId,
    handleOpenModal,
    handleOpenEditModal,
    handleCloseModal,
    handleSaveTimeSlot,
    isFormValid,
  } = useTimeSlotModal();

  const {
    isQuickModalOpen,
    quickData,
    setQuickData,
    handleOpenQuickModal,
    handleCloseQuickModal,
    handleGenerateQuickSlots,
    isQuickFormValid,
  } = useQuickSlotGenerator();

  const addTimeSlot = (slot: TimeSlot) => {
    setTimeSlots(prev => [...prev, slot]);
  };

  const updateTimeSlot = (id: string, updatedSlot: TimeSlot) => {
    setTimeSlots(prev => prev.map(slot => slot.id === id ? updatedSlot : slot));
  };

  const addTimeSlots = (slots: TimeSlot[]) => {
    setTimeSlots(prev => [...prev, ...slots]);
  };

  const handleDeleteTimeSlot = (id: string) => {
    setTimeSlots(prev => prev.filter(slot => slot.id !== id));
  };

  const handleEditTimeSlot = (slot: TimeSlot) => {
    handleOpenEditModal(slot);
  };

  return <div className="bg-success vh-100">
    <div className="container">
      <div className="row">
        <div className="col-md-6 offset-md-3 mt-5">
          <div className="card card-body">
            <h1>Plan a hangout</h1>

            <div className="mt-3">
              <label htmlFor="name">Title (optional)</label>
              <input
                id="name"
                type="text"
                className="form-control"
                placeholder="Title of the hangout or activity"
              />
            </div>

            <div className="mt-3">
              <h4>Time slots</h4>
              <div className="card card-body">
                <div className="d-flex gap-2">
                  <button
                    onClick={handleOpenModal}
                    className="btn btn-info">
                    <i className="ri-add-line" /> Add time slot
                  </button>
                  <button
                    onClick={handleOpenQuickModal}
                    className="btn btn-success">
                    <i className="ri-flashlight-line" /> Create quick slots
                  </button>
                </div>

                <TimeSlotList
                  timeSlots={timeSlots}
                  onDelete={handleDeleteTimeSlot}
                  onEdit={handleEditTimeSlot}
                />
              </div>
            </div>

            <Modal
              title={editingId ? "Edit time slot" : "Add time slot"}
              isVisible={isModalOpen}
              onClose={handleCloseModal}
              onConfirm={() => { handleSaveTimeSlot(addTimeSlot, updateTimeSlot); }}
              confirmText={editingId ? "Save" : "Add"}
              isConfirmDisabled={!isFormValid}>
              <TimeSlotForm
                date={modalData.date}
                startTime={modalData.startTime}
                endTime={modalData.endTime}
                onDateChange={(date) => {
                  setModalData(prev => ({ ...prev, date }));
                }}
                onStartTimeChange={(startTime) => {
                  setModalData(prev => ({ ...prev, startTime }));
                }}
                onEndTimeChange={(endTime) => {
                  setModalData(prev => ({ ...prev, endTime }));
                }}
              />
            </Modal>

            <Modal
              title="Create quick slots"
              isVisible={isQuickModalOpen}
              onClose={handleCloseQuickModal}
              onConfirm={() => { handleGenerateQuickSlots(addTimeSlots); }}
              confirmText="Generate"
              isConfirmDisabled={!isQuickFormValid}>
              <QuickSlotGenerator
                startDate={quickData.startDate}
                endDate={quickData.endDate}
                dailyStartTime={quickData.dailyStartTime}
                dailyEndTime={quickData.dailyEndTime}
                duration={quickData.duration}
                isOverlapping={quickData.isOverlapping}
                onStartDateChange={(startDate) => {
                  setQuickData(prev => ({ ...prev, startDate }));
                }}
                onEndDateChange={(endDate) => {
                  setQuickData(prev => ({ ...prev, endDate }));
                }}
                onDailyStartTimeChange={(dailyStartTime) => {
                  setQuickData(prev => ({ ...prev, dailyStartTime }));
                }}
                onDailyEndTimeChange={(dailyEndTime) => {
                  setQuickData(prev => ({ ...prev, dailyEndTime }));
                }}
                onDurationChange={(duration) => {
                  setQuickData(prev => ({ ...prev, duration }));
                }}
                onOverlappingChange={(isOverlapping) => {
                  setQuickData(prev => ({ ...prev, isOverlapping }));
                }}
              />
            </Modal>

            <div className="mt-3">
              <button className="btn btn-primary w-100">
                Create a new hangout
              </button>
            </div>

            <div className="mt-3">
              <Collapse title="Advanced options">
                <div className="mt-2">
                  <label htmlFor="password">Password required to join</label>
                  <input
                    id="password"
                    type="text"
                    className="form-control"
                    placeholder="Password required to join"
                  />
                </div>

                <div className="mt-3">
                  <label htmlFor="description">Description</label>
                  <input
                    id="description"
                    type="text"
                    className="form-control"
                    placeholder="Description"
                  />
                </div>

                <div className="mt-3">
                  <label htmlFor="location">Location used for weather</label>
                  <input
                    id="location"
                    type="text"
                    className="form-control"
                    placeholder="Location"
                  />
                </div>

                <div className="mt-3">
                  <label htmlFor="end-date">Accept responses until</label>
                  <input
                    id="end-date"
                    type="date"
                    className="form-control"
                    placeholder="End date"
                  />
                </div>
              </Collapse>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>;
};