import { useState } from "react";
import type { TimeSlot } from "../types";

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
    const startDate = new Date(slot.startDate);
    const endDate = new Date(slot.endDate);
    setModalData({
      date: startDate.toISOString().split("T")[0],
      startTime: startDate.toTimeString().split(" ")[0].substring(0, 5),
      endTime: endDate.toTimeString().split(" ")[0].substring(0, 5),
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
    const startDate = `${modalData.date}T${modalData.startTime}:00`;
    const endDate = `${modalData.date}T${modalData.endTime}:00`;

    if (editingId) {
      const updatedSlot: TimeSlot = {
        id: editingId,
        startDate,
        endDate,
        createdAt: new Date().toISOString(),
        updatedAt: new Date().toISOString(),
      };
      onUpdate(editingId, updatedSlot);
    } else {
      const newSlot: TimeSlot = {
        id: crypto.randomUUID(),
        startDate,
        endDate,
        createdAt: new Date().toISOString(),
        updatedAt: new Date().toISOString(),
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
