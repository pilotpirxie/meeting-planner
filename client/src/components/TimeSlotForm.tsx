import { useState } from "react";
import type { TimeSlot } from "../types";

interface TimeSlotFormProps {
  date: string;
  startTime: string;
  endTime: string;
  onDateChange: (date: string) => void;
  onStartTimeChange: (time: string) => void;
  onEndTimeChange: (time: string) => void;
}

export const useTimeSlotModal = () => {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [editingId, setEditingId] = useState<string | null>(null);
  const [modalData, setModalData] = useState({
    date: "",
    startTime: "",
    endTime: "",
  });

  const handleOpenModal = () => {
    setEditingId(null);
    setModalData({ date: "", startTime: "", endTime: "" });
    setIsModalOpen(true);
  };

  const handleOpenEditModal = (slot: TimeSlot) => {
    setEditingId(slot.id);
    setModalData({
      date: slot.date,
      startTime: slot.startTime,
      endTime: slot.endTime,
    });
    setIsModalOpen(true);
  };

  const handleCloseModal = () => {
    setIsModalOpen(false);
    setEditingId(null);
    setModalData({ date: "", startTime: "", endTime: "" });
  };

  const handleSaveTimeSlot = (
    onAdd: (slot: TimeSlot) => void,
    onUpdate: (id: string, slot: TimeSlot) => void
  ) => {
    if (editingId) {
      // Editing existing slot
      const updatedSlot: TimeSlot = {
        id: editingId,
        date: modalData.date,
        startTime: modalData.startTime,
        endTime: modalData.endTime,
      };
      onUpdate(editingId, updatedSlot);
    } else {
      // Adding new slot
      const newSlot: TimeSlot = {
        id: crypto.randomUUID(),
        date: modalData.date,
        startTime: modalData.startTime,
        endTime: modalData.endTime,
      };
      onAdd(newSlot);
    }
    handleCloseModal();
  };

  const isFormValid = modalData.date && modalData.startTime && modalData.endTime;

  return {
    isModalOpen,
    modalData,
    setModalData,
    editingId,
    handleOpenModal,
    handleOpenEditModal,
    handleCloseModal,
    handleSaveTimeSlot,
    isFormValid,
  };
};

export const TimeSlotForm = ({
  date,
  startTime,
  endTime,
  onDateChange,
  onStartTimeChange,
  onEndTimeChange,
}: TimeSlotFormProps) => {
  return (
    <>
      <div className="mt-3">
        <label htmlFor="date">Date</label>
        <input
          id="date"
          type="date"
          className="form-control"
          value={date}
          onChange={(e) => {
            onDateChange(e.target.value);
          }}
        />
      </div>

      <div className="mt-3">
        <label htmlFor="start-time">Start time</label>
        <input
          id="start-time"
          type="time"
          className="form-control"
          value={startTime}
          onChange={(e) => {
            onStartTimeChange(e.target.value);
          }}
        />
      </div>

      <div className="mt-3">
        <label htmlFor="end-time">End time</label>
        <input
          id="end-time"
          type="time"
          className="form-control"
          value={endTime}
          onChange={(e) => {
            onEndTimeChange(e.target.value);
          }}
        />
      </div>
    </>
  );
};
