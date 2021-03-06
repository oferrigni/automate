import { TestBed } from '@angular/core/testing';
import { ClientRunsSearchBarComponent } from './client-runs-search-bar.component';
import { MockComponent } from 'ng2-mock-component';
import { List } from 'immutable';
import { Chicklet } from '../../types/types';

describe('ClientRunsSearchBarComponent', () => {
  let component: ClientRunsSearchBarComponent;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [
        ClientRunsSearchBarComponent,
        MockComponent({ selector: 'chef-icon' }),
        MockComponent({ selector: 'chef-button' })
      ]
    });

    component = TestBed.createComponent(ClientRunsSearchBarComponent).componentInstance;
    component.filterTypes = [
      {
        type: 'attribute',
        text: 'Attribute'
      },
      {
        type: 'cookbook',
        text: 'Cookbook'
      },
      {
        type: 'environment',
        text: 'Environment'
      },
      {
        type: 'name',
        text: 'Node Name'
      },
      {
        type: 'platform',
        text: 'Platform'
      },
      {
        type: 'policy_group',
        text: 'Policy Group'
      },
      {
        type: 'policy_name',
        text: 'Policy Name'
      },
      {
        type: 'policy_revision',
        text: 'Policy Revision'
      },
      {
        type: 'recipe',
        text: 'Recipe'
      },
      {
        type: 'resource_name',
        text: 'Resource Name'
      },
      {
        type: 'role',
        text: 'Role'
      }
    ];
    component.visibleCategories = List<Chicklet>(component.filterTypes);
  });

  it('when clicking the category name, the selected category is removed', () => {
    component.selectedCategoryType = { type: 'attribute', text: 'Attribute' };

    component.handleCategoryClick();

    // check that the category is deselected
    expect(component.inputField.nativeElement.value).toEqual('');
    expect(component.selectedCategoryType).toEqual(undefined);
  });

  describe('when pressing arrowdown', () => {
    it('no category is selected, no input text, and nothing is highlighted: ' +
    'highlight the first category', () => {
      component.highlightedIndex = -1;

      component.handleInput('arrowdown', '');

      expect(component.highlightedIndex).toEqual(0);
      expect(component.suggestionsVisible).toEqual(true);
    });

    it('no category is selected, no input text, and the last category is highlighted: ' +
    'do not change the highlight', () => {
      component.highlightedIndex = component.visibleCategories.size - 1;

      component.handleInput('arrowdown', '');

      expect(component.highlightedIndex).toEqual(component.visibleCategories.size - 1);
      expect(component.suggestionsVisible).toEqual(true);
    });

    it('a category is selected, 3 suggestions, and none of the suggestions are highlighted: ' +
    'highlight the first suggestion', () => {
      component.selectedCategoryType = { type: 'attribute', text: 'Attribute' };
      component.highlightedIndex = -1;
      component.suggestions = List<string>(['1', '2', '3']);

      component.handleInput('arrowdown', '');

      expect(component.highlightedIndex).toEqual(0);
      expect(component.suggestionsVisible).toEqual(true);
    });

    it('a category is selected, 3 suggestions, and the second suggestions is highlighted: ' +
    'highlight the third suggestion', () => {
      component.selectedCategoryType = { type: 'attribute', text: 'Attribute' };
      component.highlightedIndex = 1;
      component.suggestions = List<string>(['1', '2', '3']);

      component.handleInput('arrowdown', '');

      expect(component.highlightedIndex).toEqual(2);
      expect(component.suggestionsVisible).toEqual(true);
    });

    it('a category is selected, 3 suggestions, and the last suggestions is highlighted: ' +
    'do not change the highlight', () => {
      component.selectedCategoryType = { type: 'attribute', text: 'Attribute' };
      component.highlightedIndex = 2;
      component.suggestions = List<string>(['1', '2', '3']);

      component.handleInput('arrowdown', '');

      expect(component.highlightedIndex).toEqual(2);
      expect(component.suggestionsVisible).toEqual(true);
    });
  });

  describe('when pressing arrowup', () => {
    it('no category is selected, no input text, and nothing is highlighted: ' +
    'do not change the highlight', () => {
      component.highlightedIndex = -1;

      component.handleInput('arrowup', '');

      expect(component.highlightedIndex).toEqual(-1);
      expect(component.suggestionsVisible).toEqual(true);
    });

    it('no category is selected, no input text, and the last category is highlighted: ' +
    'highlight the last category', () => {
      component.highlightedIndex = component.visibleCategories.size - 1;

      component.handleInput('arrowup', '');

      expect(component.highlightedIndex).toEqual(component.visibleCategories.size - 2);
      expect(component.suggestionsVisible).toEqual(true);
    });

    it('a category is selected, 3 suggestions, and the last suggestions is highlighted: ' +
    'highlight the previous suggestion', () => {
      component.selectedCategoryType = { type: 'attribute', text: 'Attribute' };
      component.suggestions = List<string>(['1', '2', '3']);
      component.highlightedIndex = component.suggestions.size - 1;

      component.handleInput('arrowup', '');

      expect(component.highlightedIndex).toEqual(component.suggestions.size - 2);
      expect(component.suggestionsVisible).toEqual(true);
    });

    it('a category is selected, 3 suggestions, and the second suggestions is highlighted: ' +
    'highlight the first suggestion', () => {
      component.selectedCategoryType = { type: 'attribute', text: 'Attribute' };
      component.highlightedIndex = 1;
      component.suggestions = List<string>(['1', '2', '3']);

      component.handleInput('arrowup', '');

      expect(component.highlightedIndex).toEqual(0);
      expect(component.suggestionsVisible).toEqual(true);
    });

    it('a category is selected, 3 suggestions, and none of the suggestions are highlighted: ' +
    'do not change the highlight', () => {
      component.selectedCategoryType = { type: 'attribute', text: 'Attribute' };
      component.highlightedIndex = -1;
      component.suggestions = List<string>(['1', '2', '3']);

      component.handleInput('arrowup', '');

      expect(component.highlightedIndex).toEqual(-1);
      expect(component.suggestionsVisible).toEqual(true);
    });
  });

  it('when pressing escape the dropdown is hidden', () => {
    component.suggestionsVisible = true;

    component.handleInput('escape', '');

    expect(component.suggestionsVisible).toEqual(false);
  });

  it('when text is typed the dropdown is displayed', () => {
    component.suggestionsVisible = false;

    component.handleInput('a', 'boba');

    expect(component.suggestionsVisible).toEqual(true);
  });

  describe('when a category is selected', () => {
    it('when backspace is pressed and the input text empty ' +
    'remove the selected category', () => {
      // Send a backspace key event when the search text is empty
      component.inputText = '';
      component.selectedCategoryType = { type: 'attribute', text: 'Attribute' };

      component.handleInput('backspace', '');

      // check that the category is deselected
      expect(component.inputField.nativeElement.value).toEqual('');
      expect(component.selectedCategoryType).toEqual(undefined);
    });

    it('text is typed: ensure a request for suggestions is made', (done) => {
      component.selectedCategoryType = { type: 'attribute', text: 'Attribute' };

      component.suggestValues.subscribe(({detail: {type: type, text: text}}) => {
        expect(text).toEqual('boba');
        expect(type).toEqual('attribute');
        done();
      });

      component.handleInput('a', 'boba');
    });

    it('when text is typed the dropdown is displayed', () => {
      component.selectedCategoryType = { type: 'attribute', text: 'Attribute' };
      component.suggestionsVisible = false;

      component.handleInput('a', 'boba');

      expect(component.suggestionsVisible).toEqual(true);
    });

    describe('enter is typed', () => {
      it('input text is empty, ' +
      'and a suggested item is highlighted: select the highlighted item', (done) => {
        component.selectedCategoryType = { type: 'attribute', text: 'Attribute' };
        component.suggestions = List<string>(['1', '2', '3']);
        component.highlightedIndex = 1;

        component.itemSelected.subscribe(({detail: {type: type, text: text}}) => {
          expect(text).toEqual('2');
          expect(type).toEqual('attribute');
          done();
        });

        component.handleInput('enter', '');
      });

      it('no items are highlighted, and the input text matching a suggestion is available: ' +
      'select the matching suggestion item', (done) => {
        component.selectedCategoryType = { type: 'attribute', text: 'Attribute' };
        component.suggestions = List<string>(['1', '2', '3']);
        component.highlightedIndex = -1;

        component.itemSelected.subscribe(({detail: {type: type, text: text}}) => {
          expect(text).toEqual('2');
          expect(type).toEqual('attribute');
          done();
        });

        component.handleInput('enter', '2');
      });

      it('text matching a suggestion is not available: do not select an item', () => {
        jasmine.clock().uninstall();
        jasmine.clock().install();
        component.selectedCategoryType = { type: 'attribute', text: 'Attribute' };
        component.suggestions = List<string>(['1', '2', 'bbbb']);
        component.highlightedIndex = -1;

        let isItemSelectedCalled = false;

        component.itemSelected.subscribe((_result: any) => {
          fail('Should not select an item');
          isItemSelectedCalled = true;
        });

        component.handleInput('enter', 'bbb');

        jasmine.clock().tick(1000);
        jasmine.clock().uninstall();

        expect(isItemSelectedCalled).toBe(false);
      });

      describe('input text contains wildcards', () => {
        it('no items are highlighted: select the wildcard * item', (done) => {
          component.selectedCategoryType = { type: 'name', text: 'Node Name' };
          component.suggestions = List<string>(['chef-load-1', 'chef-load-2', 'chef-load-3']);
          component.highlightedIndex = -1;

          component.itemSelected.subscribe(({detail: {type: type, text: text}}) => {
            expect(text).toEqual('chef-load-*');
            expect(type).toEqual('name');
            done();
          });

          component.handleInput('enter', 'chef-load-*');
        });

        it('no items are highlighted: select the wildcard ? item', (done) => {
          component.selectedCategoryType = { type: 'name', text: 'Node Name' };
          component.suggestions = List<string>(['chef-1-nodes', 'chef-2-nodes', 'chef-3-nodes']);
          component.highlightedIndex = -1;

          component.itemSelected.subscribe(({detail: {type: type, text: text}}) => {
            expect(text).toEqual('chef-?-nodes');
            expect(type).toEqual('name');
            done();
          });

          component.handleInput('enter', 'chef-?-nodes');
        });
      });
    });
  });

  describe('when a category is not selected and enter is typed', () => {
    it('there are more than one visible categories and none are highlighted: ' +
    'do not select a category', () => {
      component.highlightedIndex = -1;
      component.selectedCategoryType = undefined;

      // Set the text to name, which will match three categories
      component.handleInput('e', 'name');

      expect(component.visibleCategories.size).toEqual(3);

      // pressing enter with 'name' in the text field
      component.handleInput('enter', 'name');

      // A category should not be selected.
      expect(component.selectedCategoryType).toEqual(undefined);
    });

    it('there are only one visible categories: select a category', () => {
      component.highlightedIndex = -1;
      component.selectedCategoryType = undefined;

      // Set the text to node, which will match only one category 'Node Name'
      component.handleInput('e', 'node');

      expect(component.visibleCategories.size).toEqual(1);

      // pressing enter with 'node' in the text field
      component.handleInput('enter', 'node');

      // A category should be selected.
      expect(component.selectedCategoryType).toEqual({ type: 'name', text: 'Node Name'});
    });

    it('there are more than one visible categories, and one of the categories is highlighted: ' +
      'select the highlighted category', () => {
      component.selectedCategoryType = undefined;
      component.highlightedIndex = 0;

      component.handleInput('enter', '');

      // A category should be selected.
      expect(component.selectedCategoryType).toEqual({ type: 'attribute', text: 'Attribute'});
    });
  });
});
