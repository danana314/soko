/**
 * Schedule Manager - Alpine.js component for trip schedule functionality
 * Handles bulk selection, auto-save, and visual feedback for schedule changes
 */
document.addEventListener('alpine:init', () => {
    Alpine.data('scheduleManager', () => ({
        saveTimeout: null,
        saving: false,
        savedRecently: false,
        
        /**
         * Toggle all checkboxes for a specific user (column)
         * @param {string} userId - The ID of the user
         */
        toggleUserColumn(userId) {
            // Find all checkboxes for this user
            const checkboxes = document.querySelectorAll(`input[data-user-id="${userId}"]`);
            
            // Check if all are currently checked
            const allChecked = Array.from(checkboxes).every(cb => cb.checked);
            
            // Toggle all to opposite state
            checkboxes.forEach(checkbox => {
                checkbox.checked = !allChecked;
            });
            
            // Auto-save after bulk change
            this.scheduleChanged();
        },
        
        /**
         * Toggle all checkboxes for a specific date (row)
         * @param {string} date - The formatted date string
         */
        toggleDateRow(date) {
            // Find all checkboxes for this date
            const checkboxes = document.querySelectorAll(`input[data-date="${date}"]`);
            
            // Check if all are currently checked
            const allChecked = Array.from(checkboxes).every(cb => cb.checked);
            
            // Toggle all to opposite state
            checkboxes.forEach(checkbox => {
                checkbox.checked = !allChecked;
            });
            
            // Auto-save after bulk change
            this.scheduleChanged();
        },
        
        /**
         * Handle schedule change events with debouncing
         * Delays save by 500ms to prevent excessive requests during rapid changes
         */
        scheduleChanged() {
            // Clear existing timeout
            if (this.saveTimeout) {
                clearTimeout(this.saveTimeout);
            }
            
            // Set new timeout for debounced save
            this.saveTimeout = setTimeout(() => {
                this.saveSchedule();
            }, 500);
        },
        
        /**
         * Save the current schedule state to the server
         * Uses fetch API to maintain current tab state and provide visual feedback
         */
        saveSchedule() {
            this.saving = true;
            
            const form = document.querySelector('form[hx-post*="/s"]');
            if (!form) {
                this.saving = false;
                return;
            }
            
            // Get all checked checkboxes and build form data manually
            const formData = new FormData();
            const checkedBoxes = form.querySelectorAll('input[type="checkbox"]:checked');
            checkedBoxes.forEach(checkbox => {
                formData.append(checkbox.name, checkbox.value || 'on');
            });
            
            // Save to server without updating DOM to preserve tab state
            fetch(form.getAttribute('hx-post'), {
                method: 'POST',
                body: new URLSearchParams(Object.fromEntries(formData))
            }).then(response => {
                if (!response.ok) {
                    throw new Error(`HTTP ${response.status}`);
                }
                
                this.saving = false;
                this.savedRecently = true;
                
                // Hide success indicator after 2 seconds
                setTimeout(() => {
                    this.savedRecently = false;
                }, 2000);
            }).catch(error => {
                console.error('Schedule save failed:', error);
                this.saving = false;
                
                // TODO: Add error state handling
                // Could show error message to user here
            });
        }
    }));
});